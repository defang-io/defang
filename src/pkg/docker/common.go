package docker

import (
	"context"
	"errors"

	"github.com/defang-io/defang/src/pkg/types"
	"github.com/docker/docker/client"
)

type ContainerID = types.TaskID

type Docker struct {
	*client.Client

	image    string
	memory   uint64
	platform string
}

func New() *Docker {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return &Docker{
		Client: cli,
	}
}

var _ types.Driver = (*Docker)(nil)

func (Docker) SetVpcID(vpcId string) error {
	if vpcId != "" {
		return errors.New("docker does not support specifying VPC ID")
	}
	return nil
}

func (Docker) PutSecret(ctx context.Context, name, value string) error {
	return errors.New("docker does not support secrets")
}

func (Docker) ListSecrets(ctx context.Context) ([]string, error) {
	return nil, errors.New("docker does not support secrets")
}
