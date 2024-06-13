#!/usr/bin/env node
import AdmZip from "adm-zip";
import axios from "axios";
import * as child_process from "child_process";
import { promises as fsPromise } from "fs";
import * as os from "os";
import * as path from "path";
import * as tar from "tar";
import { promisify } from "util";

const EXECUTABLE = "defang";
const URL_LATEST_RELEASE =
  "https://api.github.com/repos/DefangLabs/defang/releases/latest";
const HTTP_STATUS_OK = 200;

const exec = promisify(child_process.exec);

async function getLatestVersion(): Promise<string> {
  const response = await axios.get(URL_LATEST_RELEASE);
  if (response.status !== HTTP_STATUS_OK) {
    throw new Error(
      `Failed to get latest version from GitHub. Status code: ${response.status}`
    );
  }

  return response.data.tag_name.replace("v", "").trim();
}

async function downloadAppArchive(
  version: string,
  archiveFilename: string,
  outputPath: string
): Promise<string | null> {
  const repo = "DefangLabs/defang";
  const downloadUrl = `https://github.com/${repo}/releases/download/v${version}/${archiveFilename}`;
  const downloadTargetFile = path.join(outputPath, archiveFilename);

  return await downloadFile(downloadUrl, downloadTargetFile);
}

async function downloadFile(
  downloadUrl: string,
  downloadTargetFile: string
): Promise<string | null> {
  try {
    const response = await axios.get(downloadUrl, {
      responseType: "arraybuffer",
      headers: {
        "Content-Type": "application/octet-stream",
      },
    });

    if (response?.data === undefined) {
      throw new Error(
        `Failed to download ${downloadUrl}. No data in response.`
      );
    }

    // write data to file, will overwrite if file already exists
    await fsPromise.writeFile(downloadTargetFile, response.data);

    return downloadTargetFile;
  } catch (error) {
    console.error(error);

    // something went wrong, clean up by deleting the downloaded file if it exists
    await fsPromise.unlink(downloadTargetFile);
    return null;
  }
}

async function extractArchive(
  archiveFilePath: string,
  outputPath: string
): Promise<boolean> {
  let result = false;

  const ext = path.extname(archiveFilePath).toLocaleLowerCase();
  switch (ext) {
    case ".zip":
      result = await extractZip(archiveFilePath, outputPath);
      break;
    case ".gz":
      result = extractTarGz(archiveFilePath, outputPath);
      break;
    default:
      throw new Error(`Unsupported archive extension: ${ext}`);
  }

  return result;
}

async function extractZip(
  zipPath: string,
  outputPath: string
): Promise<boolean> {
  try {
    const zip = new AdmZip(zipPath);
    const result = zip.extractEntryTo(EXECUTABLE, outputPath, true, true);
    await fsPromise.chmod(path.join(outputPath, EXECUTABLE), 755);
    return result;
  } catch (error) {
    console.error(`An error occurred during zip extraction: ${error}`);
    return false;
  }
}

function extractTarGz(tarGzFilePath: string, outputPath: string): boolean {
  try {
    tar.extract(
      {
        cwd: outputPath,
        file: tarGzFilePath,
        sync: true,
        strict: true,
      },
      [EXECUTABLE]
    );
    return true;
  } catch (error) {
    console.error(`An error occurred during tar.gz extraction: ${error}`);
    return false;
  }
}

async function deleteArchive(archiveFilePath: string): Promise<void> {
  await fsPromise.unlink(archiveFilePath);
}

async function getVersion(filename: string): Promise<string> {
  const data = await fsPromise.readFile(filename, "utf8");
  const pkg = JSON.parse(data);
  return pkg.version;
}

function getAppArchiveFilename(
  version: string,
  platform: string,
  arch: string
): string {
  let compression = "zip";
  switch (platform) {
    case "windows":
      platform = "windows";
      break;
    case "linux":
      platform = "linux";
      compression = "tar.gz";
      break;
    case "darwin":
      platform = "macOS";
      break;
    default:
      throw new Error(`Unsupported operating system: ${platform}`);
  }

  switch (arch) {
    case "x64":
      arch = "amd64";
      break;
    case "arm64":
      arch = "arm64";
      break;
    default:
      throw new Error(`Unsupported architecture: ${arch}`);
  }

  if (platform === "macOS") {
    return `defang_${version}_${platform}.${compression}`;
  }
  return `defang_${version}_${platform}_${arch}.${compression}`;
}

async function install(version: string, saveDirectory: string) {
  try {
    console.log(`Getting latest defang cli`);

    // download the latest version of defang cli
    const filename = getAppArchiveFilename(version, os.platform(), os.arch());
    const archiveFile = await downloadAppArchive(
      version,
      filename,
      saveDirectory
    );

    if (archiveFile == null || archiveFile.length === 0) {
      throw new Error(`Failed to download ${filename}`);
    }

    // Because the releases are compressed tar.gz or .zip we need to
    // uncompress them to the ./bin directory in the package in node_modules.
    const result = await extractArchive(archiveFile, saveDirectory);
    if (result === false) {
      throw new Error(`Failed to install binaries!`);
    }

    // Delete the downloaded archive since we have successfully downloaded
    // and uncompressed it.
    await deleteArchive(archiveFile);
  } catch (error) {
    console.error(error);
  }
}

function getPathToExecutable(): string | null {
  let extension = "";
  if (["win32", "cygwin"].includes(process.platform)) {
    extension = ".exe";
  }

  const executablePath = path.join(__dirname, `${EXECUTABLE}${extension}`);
  try {
    return require.resolve(executablePath);
  } catch (e) {
    return null;
  }
}

function extractCLIVersions(versionInfo: string): {
  defangCLI: string;
  latestCLI: string;
} {
  // parse the CLI version info
  // e.g.
  // Defang CLI:    v0.5.24
  // Latest CLI:    v0.5.24
  // Defang Fabric: v0.5.0-643-abcdef012
  //

  const versionRegex = /\d+\.\d+\.\d+/g;
  const matches = versionInfo.match(versionRegex);

  if (matches != null && matches.length >= 2) {
    return {
      defangCLI: matches[0],
      latestCLI: matches[1],
    };
  } else {
    throw new Error("Could not extract CLI versions from the output.");
  }
}

type VersionInfo = {
  current: string | null;
  latest: string | null;
};

async function getVersionInfo(): Promise<VersionInfo> {
  let result: VersionInfo = { current: null, latest: null };
  try {
    const execPath = getPathToExecutable();

    if (!execPath) {
      // there is no executable, so we can't get the version info from the CLI
      const latestVersion = await getLatestVersion();

      return { current: null, latest: latestVersion };
    }

    // Exec output contains both stderr and stdout outputs
    const versionInfo = await exec(execPath + " version");

    const verInfo = extractCLIVersions(versionInfo.stdout);
    result.current = verInfo.defangCLI;
    result.latest = verInfo.latestCLI;
  } catch (error) {
    console.error(error);
  }

  return result;
}

// js wrapper to use by npx or npm exec, this will call the defang binary with
// the arguments passed to the npx line. NPM installer will create a symlink
// in the user PATH to the cli.js to execute. The symlink will name the same as
// the package name i.e. defang.
async function run(): Promise<void> {
  try {
    const { current, latest }: VersionInfo = await getVersionInfo();

    // get the latest version of defang cli if not already installed
    if (latest != null && (current == null || current != latest)) {
      await install(latest, __dirname);
    }

    // execute the defang binary with the arguments passed to the npx line.
    const args = process.argv.slice(2);
    const pathToExec = getPathToExecutable();
    if (!pathToExec) {
      throw new Error("Could not find the defang executable.");
    }

    const processResult = child_process.spawnSync(pathToExec, args, {
      stdio: "inherit",
    });

    // if there was an error, print it to the console.
    processResult.error && console.error(processResult.error);
    process.exitCode = processResult.status ?? 1;
  } catch (error) {
    console.error(error);
    process.exitCode = 2;
  }
}

run();
