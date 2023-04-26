//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pterm/pterm"
)

const (
	// PortMapping is the port mapping for the mkdocs server running locally.
	portMapping = "3001:8000"

	// MkdocsDockerImage is the docker image to use for running mkdocs.
	// mkdocsDockerImage = "squidfunk/mkdocs-material"

	// MkdocsDockerTag is the docker tag to use for running mkdocs.
	// mkdocsDockerTag = "latest"

	// LocalDockerImageName is the Docker image to run commands against after building.
	localDockerImageName = "dev.local/mkdocs:latest"

	// DockerFileName is the Dockerfile to use for building the image for running locally.
	dockerFileName = "Dockerfile.mkdocs"

	// LocalContainerName is the name of the container to run locally, and can be stopped too.
	localContainerName = "mkdocs"
)

// QualifiedDockerImage is the fully qualified docker image to use for running mkdocs.
// var _qualifiedDockerImage = fmt.Sprintf("%s:%s", mkdocsDockerImage, mkdocsDockerTag)

// invokeMKDocs runs the mkdocs serve command via docker.
func invokeMKDocs(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	invokeArgs := []string{
		"run",
		"--name", localContainerName,
		"--rm",
		"-it",
		"-p",
		portMapping,
		"-v",
		fmt.Sprintf("%s:/docs", wd),
		localDockerImageName,
	}
	invokeArgs = append(invokeArgs, args...)
	if mg.Verbose() {
		invokeArgs = append(invokeArgs, "--verbose")
	}
	pterm.Success.Printfln("server on host via http://localhost:%s", strings.Split(portMapping, ":")[0])
	return sh.RunV("docker", invokeArgs...)
}

// 🌐 Run mkdocs serve via Docker.
func Serve() error {
	mg.Deps(Build)
	return invokeMKDocs("serve", "--dev-addr", "0.0.0.0:8000") // Required to pass 0.0.0.0 to ensure that the server is accessible from the host when running in docker.
}

// 🔨 Run docker build.
//
// This is required as custom plugins are installed in docker image.
func Build() error {
	return sh.Run("docker", "build", "-t", localDockerImageName, "-f", dockerFileName, ".")
}

// 🛑 Stop the mkdocs dockerized container.
func Stop() {
	if err := sh.Run("docker", "stop", localContainerName); err != nil {
		pterm.Warning.Printfln("container %s not found", localContainerName)
	}
}
