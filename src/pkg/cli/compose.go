package cli

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
	pb "github.com/defang-io/defang/src/protos/io/defang/v1"
	"github.com/defang-io/defang/src/protos/io/defang/v1/defangv1connect"
	"github.com/moby/patternmatcher"
	"github.com/moby/patternmatcher/ignorefile"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	MiB                 = 1024 * 1024
	sourceDateEpoch     = 315532800 // 1980-01-01, same as nix-shell
	defaultDockerIgnore = `# Default .dockerignore file for Defang
**/.DS_Store
**/.direnv
**/.envrc
**/.git
**/.github
**/.idea
**/.vscode
**/__pycache__
**/compose.yaml
**/compose.yml
**/defang.exe
**/docker-compose.yml
**/docker-compose.yaml
**/node_modules
**/Thumbs.db
# Ignore our own binary, but only in the root to avoid ignoring subfolders
defang`
)

var (
	nonAlphanumeric = regexp.MustCompile(`[^a-zA-Z0-9]+`)
)

type ComposeError struct {
	error
}

func (e ComposeError) Unwrap() error {
	return e.error
}

func NormalizeServiceName(s string) string {
	return nonAlphanumeric.ReplaceAllLiteralString(strings.ToLower(s), "-")
}

func resolveEnv(k string) *string {
	// TODO: per spec, if the value is nil, then the value is taken from an interactive prompt
	v, ok := os.LookupEnv(k)
	if !ok {
		logrus.Warnf("environment variable not found: %q", k)
		// If the value could not be resolved, it should be removed
		return nil
	}
	return &v
}

func convertPlatform(platform string) pb.Platform {
	switch platform {
	default:
		logrus.Warnf("Unsupported platform: %q (assuming linux)", platform)
		fallthrough
	case "", "linux":
		return pb.Platform_LINUX_ANY
	case "linux/amd64":
		return pb.Platform_LINUX_AMD64
	case "linux/arm64", "linux/arm64/v8", "linux/arm64/v7", "linux/arm64/v6":
		return pb.Platform_LINUX_ARM64
	}
}

func loadDockerCompose(filePath, projectName string) (*types.Project, error) {
	Debug(" - Loading compose file", filePath, "for project", projectName)
	// Compose-go uses the logrus logger, so we need to configure it to be more like our own logger
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, DisableColors: !DoColor, DisableLevelTruncation: true})
	project, err := loader.Load(types.ConfigDetails{
		WorkingDir:  filepath.Dir(filePath),
		ConfigFiles: []types.ConfigFile{{Filename: filePath}},
		Environment: map[string]string{}, // TODO: support environment variables?
	}, loader.WithDiscardEnvFiles, func(o *loader.Options) {
		o.SetProjectName(strings.ToLower(projectName), projectName != "") // normalize to lowercase
		o.SkipConsistencyCheck = true                                     // TODO: check fails if secrets are used but top-level 'secrets:' is missing
	})
	if err != nil {
		return nil, err
	}

	if DoVerbose {
		b, _ := yaml.Marshal(project)
		fmt.Println(string(b))
	}
	return project, nil
}

func getRemoteBuildContext(ctx context.Context, client defangv1connect.FabricControllerClient, name string, build *types.BuildConfig, force bool) (string, error) {
	root, err := filepath.Abs(build.Context)
	if err != nil {
		return "", fmt.Errorf("invalid build context: %w", err)
	}

	Info(" * Compressing build context for", name, "at", root)
	buffer, err := createTarball(ctx, build.Context, build.Dockerfile)
	if err != nil {
		return "", err
	}

	var digest string
	if !force {
		// Calculate the digest of the tarball and pass it to the fabric controller (to avoid building the same image twice)
		sha := sha256.Sum256(buffer.Bytes())
		digest = "sha256-" + base64.StdEncoding.EncodeToString(sha[:]) // same as Nix
		Debug(" - Digest:", digest)
	}

	if DoDryRun {
		return root, nil
	}

	Info(" * Uploading build context for", name)
	return uploadTarball(ctx, client, buffer, digest)
}

func convertPort(port types.ServicePortConfig) (*pb.Port, error) {
	if port.Target < 1 || port.Target > 32767 {
		return nil, fmt.Errorf("port target must be an integer between 1 and 32767: %v", port.Target)
	}
	if port.HostIP != "" {
		return nil, errors.New("port host_ip is not supported")
	}
	if port.Published != "" && port.Published != strconv.FormatUint(uint64(port.Target), 10) {
		return nil, fmt.Errorf("port published must be empty or equal to target: %v", port.Published)
	}

	pbPort := &pb.Port{
		// Mode      string `yaml:",omitempty" json:"mode,omitempty"`
		// HostIP    string `mapstructure:"host_ip" yaml:"host_ip,omitempty" json:"host_ip,omitempty"`
		// Published string `yaml:",omitempty" json:"published,omitempty"`
		// Protocol  string `yaml:",omitempty" json:"protocol,omitempty"`
		Target: port.Target,
	}

	switch port.Protocol {
	case "":
		pbPort.Protocol = pb.Protocol_ANY // defaults to HTTP in CD
	case "tcp":
		pbPort.Protocol = pb.Protocol_TCP
	case "udp":
		pbPort.Protocol = pb.Protocol_UDP
	case "http": // TODO: not per spec
		pbPort.Protocol = pb.Protocol_HTTP
	case "http2": // TODO: not per spec
		pbPort.Protocol = pb.Protocol_HTTP2
	case "grpc": // TODO: not per spec
		pbPort.Protocol = pb.Protocol_GRPC
	default:
		return nil, fmt.Errorf("port protocol not one of [tcp udp http http2 grpc]: %v", port.Protocol)
	}

	logrus := logrus.WithField("target", port.Target)

	switch port.Mode {
	case "":
		logrus.Warn("No port mode was specified; assuming 'host' (add 'mode' to silence)")
		fallthrough
	case "host":
		pbPort.Mode = pb.Mode_HOST
	case "ingress":
		// This code is unnecessarily complex because compose-go silently converts short syntax to ingress+tcp
		if port.Published != "" {
			logrus.Warn("Published ports are not supported in ingress mode; assuming 'host' (add 'mode' to silence)")
			break
		}
		pbPort.Mode = pb.Mode_INGRESS
		if pbPort.Protocol == pb.Protocol_TCP || pbPort.Protocol == pb.Protocol_UDP {
			logrus.Warn("TCP ingress is not supported; assuming HTTP")
			pbPort.Protocol = pb.Protocol_HTTP
		}
	default:
		return nil, fmt.Errorf("port mode not one of [host ingress]: %v", port.Mode)
	}
	return pbPort, nil
}

func convertPorts(ports []types.ServicePortConfig) ([]*pb.Port, error) {
	var pbports []*pb.Port
	for _, port := range ports {
		pbPort, err := convertPort(port)
		if err != nil {
			return nil, err
		}
		pbports = append(pbports, pbPort)
	}
	return pbports, nil
}

func uploadTarball(ctx context.Context, client defangv1connect.FabricControllerClient, body *bytes.Buffer, digest string) (string, error) {
	// Upload the tarball to the fabric controller storage TODO: use a streaming API
	ureq := &pb.UploadURLRequest{Digest: digest}
	res, err := client.CreateUploadURL(ctx, connect.NewRequest(ureq))
	if err != nil {
		return "", err
	}

	if DoDryRun {
		return "", errors.New("dry run")
	}

	// Do an HTTP PUT to the generated URL
	req, err := http.NewRequestWithContext(ctx, "PUT", res.Msg.Url, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP PUT failed with status code %v", resp.Status)
	}

	// Remove query params from URL
	url, err := url.Parse(res.Msg.Url)
	if err != nil {
		return "", err
	}
	url.RawQuery = ""
	return url.String(), nil
}

type contextAwareWriter struct {
	ctx context.Context
	io.WriteCloser
}

func (cw contextAwareWriter) Write(p []byte) (n int, err error) {
	select {
	case <-cw.ctx.Done(): // Detect context cancelation
		return 0, cw.ctx.Err()
	default:
		return cw.WriteCloser.Write(p)
	}
}

func createTarball(ctx context.Context, root, dockerfile string) (*bytes.Buffer, error) {
	foundDockerfile := false
	if dockerfile == "" {
		dockerfile = "Dockerfile"
	} else {
		dockerfile = filepath.Clean(dockerfile)
	}

	// A Dockerfile-specific ignore-file takes precedence over the .dockerignore file at the root of the build context if both exist.
	var reader io.ReadCloser
	var err error
	reader, err = os.Open(filepath.Join(root, dockerfile+".dockerignore"))
	if err != nil {
		reader, err = os.Open(filepath.Join(root, ".dockerignore"))
		if err != nil {
			Debug(" - No .dockerignore file found; using defaults")
			reader = io.NopCloser(strings.NewReader(defaultDockerIgnore))
		} else {
			Debug(" - Reading .dockerignore file")
		}
	} else {
		Debug(" - Reading", dockerfile+".dockerignore file")
	}
	patterns, err := ignorefile.ReadAll(reader) // handles comments and empty lines
	if reader != nil {
		reader.Close()
	}
	if err != nil {
		return nil, err
	}
	pm, err := patternmatcher.New(patterns)
	if err != nil {
		return nil, err
	}

	// TODO: use io.Pipe and do proper streaming (instead of buffering everything in memory)
	fileCount := 0
	var buf bytes.Buffer
	gzipWriter := &contextAwareWriter{ctx, gzip.NewWriter(&buf)}
	tarWriter := tar.NewWriter(gzipWriter)

	err = filepath.WalkDir(root, func(path string, de os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Don't include the root directory itself in the tarball
		if path == root {
			return nil
		}

		// Make sure the path is relative to the root
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Ignore files using the dockerignore patternmatcher
		baseName := filepath.ToSlash(relPath)
		ignore, err := pm.MatchesOrParentMatches(baseName)
		if err != nil {
			return err
		}
		if ignore {
			Debug(" - Ignoring", relPath)
			if de.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		Debug(" - Adding", baseName)

		info, err := de.Info()
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Make reproducible; WalkDir walks files in lexical order.
		header.ModTime = time.Unix(sourceDateEpoch, 0)
		header.Gid = 0
		header.Uid = 0
		header.Name = baseName
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		if !foundDockerfile && dockerfile == relPath {
			foundDockerfile = true
		}

		fileCount++
		if fileCount == 11 {
			Warn(" ! The build context contains more than 10 files; press Ctrl+C if this is unexpected.")
		}

		_, err = io.Copy(tarWriter, file)
		if buf.Len() > 10*MiB {
			return fmt.Errorf("build context is too large; this beta version is limited to 10MiB")
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	// Close the tar and gzip writers before returning the buffer
	if err = tarWriter.Close(); err != nil {
		return nil, err
	}

	if err = gzipWriter.Close(); err != nil {
		return nil, err
	}

	if !foundDockerfile {
		return nil, fmt.Errorf("the specified dockerfile could not be read: %q", dockerfile)
	}

	return &buf, nil
}
