//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pterm/pterm"
)

// Mkdocs is the namespace for running mkdocs commands.
type Mkdocs mg.Namespace

// Docker is the namespace for running docker commands.
type Docker mg.Namespace

const (
	// PortMapping is the port mapping for the mkdocs server running locally.
	portMapping = "3001:8000"

	// MkdocsDockerImage is the docker image to use for running mkdocs.
	mkdocsDockerImage = "squidfunk/mkdocs-material"

	// MkdocsDockerTag is the docker tag to use for running mkdocs.
	mkdocsDockerTag = "latest"

	// LocalDockerImageName is the Docker image to run commands against after building.
	localDockerImageName = "dev.local/mkdocs:latest"

	// DockerFileName is the Dockerfile to use for building the image for running locally.
	dockerFileName = "Dockerfile.mkdocs"

	// LocalContainerName is the name of the container to run locally, and can be stopped too.
	localContainerName = "mkdocs"

	// ArtifactsDirectory is the directory to store artifacts in.
	ArtifactsDirectory = ".artifacts"

	// PermissionReadWriteOwner is the permission to use for the artifacts directory.
	PermissionReadWriteOwner = 0o700
)

const (
	// mkdocsConfigFile is the name of the mkdocs config file.
	mkdocsConfigFile = "mkdocs.yml"
)

// QualifiedDockerImage is the fully qualified docker image to use for running mkdocs.
// var _qualifiedDockerImage = fmt.Sprintf("%s:%s", mkdocsDockerImage, mkdocsDockerTag)

// invokeMKDocs runs the mkdocs serve command via docker.
func invokeMKDocs(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get working directory: %w", err)
	}
	invokeArgs := []string{
		"run",
		"--name", localContainerName,
		"--rm",
		"-it",
		"-p", portMapping,
		"-v", fmt.Sprintf("%s:/docs", wd),
	}
	invokeArgs = append(invokeArgs, localDockerImageName)
	invokeArgs = append(invokeArgs, args...)

	if mg.Verbose() {
		invokeArgs = append(invokeArgs, "--verbose")
	}
	pterm.Success.Printfln("server on host via http://localhost:%s", strings.Split(portMapping, ":")[0])
	return sh.RunV("docker", invokeArgs...)
}

// 🌐 Run mkdocs serve via Docker.
func (Mkdocs) Serve() error {
	mg.Deps(Docker{}.Build)
	return invokeMKDocs("serve", "--dev-addr", "0.0.0.0:8000") // Required to pass 0.0.0.0 to ensure that the server is accessible from the host when running in docker.
}

// 🔨 Run docker build.
//
// This is required as custom plugins are installed in docker image.
func (Docker) Build() error {
	start := time.Now()
	defer func() {
		pterm.Success.Printfln("Build() %s", humanize.RelTime(start, time.Now(), "", ""))
	}()
	return sh.Run("docker", "build", "-t", localDockerImageName, "-f", dockerFileName, ".")
}

// 🛑 Stop the mkdocs dockerized container.
func (Docker) Stop() {
	if err := sh.Run("docker", "stop", localContainerName); err != nil {
		pterm.Warning.Printfln("container %s not found", localContainerName)
	}
}

// Run mkdocs commands contained in docker
func (Mkdocs) Build() error {
	start := time.Now()
	defer func() {
		pterm.Success.Printfln("(Mkdocs) Build() %s", humanize.RelTime(start, time.Now(), "", ""))
	}()

	mg.Deps(Docker{}.Build)
	if err := os.MkdirAll(ArtifactsDirectory, PermissionReadWriteOwner); err != nil {
		return fmt.Errorf("unable to create the artifact directory at %s: %w", ArtifactsDirectory, err)
	}
	return invokeMKDocs(
		"build",
		"--config-file", mkdocsConfigFile,
		"--site-dir", filepath.Join(ArtifactsDirectory, "_site"),
	)
}

// Run mkdocs commands contained in docker
func (Mkdocs) GHDeploy() error {
	start := time.Now()
	defer func() {
		pterm.Success.Printfln("(Mkdocs) GHDeploy() %s", humanize.RelTime(start, time.Now(), "", ""))
	}()

	mg.Deps(Mkdocs{}.Build)
	return invokeMKDocs("gh-deploy", "--force")
}

// Pull pulls the squidfunk/mkdocs-material:latest Docker image
func (Docker) Pull() error {
	return sh.RunV("docker", "pull", fmt.Sprintf("%s:%s", mkdocsDockerImage, mkdocsDockerTag))
}
