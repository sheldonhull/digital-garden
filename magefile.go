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
	// portMapping is the port mapping for the mkdocs server running locally
	portMapping = "3001:8000"

	// mkdocsDockerImage is the docker image to use for running mkdocs
	mkdocsDockerImage = "squidfunk/mkdocs-material"

	// mkdocsDockerTag is the docker tag to use for running mkdocs
	mkdocsDockerTag = "latest"
)

var (
	// qualifiedDockerImage is the fully qualified docker image to use for running mkdocs
	_qualifiedDockerImage = fmt.Sprintf("%s:%s", mkdocsDockerImage, mkdocsDockerTag)
)

// run the mkdocs server via docker
func invokeMKDocs(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	invokeArgs := []string{"run", "--rm", "-it", "-p", portMapping, "-v", fmt.Sprintf("%s:/docs", wd), _qualifiedDockerImage}
	invokeArgs = append(invokeArgs, args...)
	if mg.Verbose() {
		invokeArgs = append(invokeArgs, "--verbose")
	}
	pterm.Success.Printfln("server on host via http://localhost:%s", strings.Split(portMapping, ":")[0])
	return sh.RunV("docker", invokeArgs...)
}

// Run mkdocs server
func Serve() error {
	return invokeMKDocs("serve")

}
