//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Run mkdocs server
func Serve() error {
	return sh.RunV("mkdocs", "serve")

}
