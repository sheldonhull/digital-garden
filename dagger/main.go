// A generated module for Garden functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
	"path/filepath"

	"dagger/garden/internal/dagger"
)

// MARK: constants & vars

// Garden is a struct that contains all the functions for the garden module.
type Garden struct{}

const (
	// artifactsDirectory is the directory where the artifacts are stored and ignored by git.
	artifactsDirectory = ".artifacts"
)

const (
	mkdocsBaseImage = "squidfunk/mkdocs-material"
	mkdocsTag       = "latest"
)

const (
	localRepositoryName   = "dev.local"
	localImageName        = "mkdocs"
	localImageTag         = "latest"
	localRequirementsFile = "requirements.txt"

	// daggerInternalMkdocsWorkingDirectory is the working directory for the mkdocs container, internal to the container, not the host.
	daggerInternalMkdocsWorkingDirectory = "/docs"
	// permissionUserReadWriteExecute is the permissions for the artifact directory.
	permissionUserReadWriteExecute = 0o0700

	// localHostAddressToUse is the address to use when binding to the host.
	// Using 0.0.0.0 will result in macOS prompting for loopback network security permissions.
	localHostAddressToUse = "127.0.0.1"

	// internalLocalHostAddressToUse is so that the mapping internal can be resolved by docker.
	internalLocalHostAddressToUse = "0.0.0.0"
	localHostPort                 = 8000
)

var (
	qualifiedImageName                = mkdocsBaseImage + ":" + mkdocsTag
	qualifiedLocalImageName           = localRepositoryName + "/" + localImageName + ":" + localImageTag
	qualifiedInternalRequirementsFile = filepath.Join(daggerInternalMkdocsWorkingDirectory, localRequirementsFile)
	qualifiedInternalArtifactsSiteDir = filepath.Join(daggerInternalMkdocsWorkingDirectory, ".artifacts", "site")
	qualifiedLocalArtifactsSiteDir    = filepath.Join(artifactsDirectory, "_site")
)

// MARK: Dagger Functions

// MkdocsBaseContainer creates the docker image base container that other actions can be run from.
func (m *Garden) MkdocsBaseContainer(ctx context.Context, dir *Directory) *Container {
	packages := []string{
		"mkdocs-glightbox",
		"mkdocs-rss-plugin",
		"mkdocs-autolinks-plugin",
		"mkdocs-git-revision-date-localized-plugin",
		"mkdocs-exclude",
		"mkdocs-git-authors-plugin",
		"mkdocs-swagger-ui-tag",
		"mkdocs-glightbox",
		"markdown-callouts",
		"mkdocs-awesome-pages-plugin",
	}
	fmt.Printf("package listing: %v", packages)
	ctr := dag.Container().From(qualifiedImageName).
		WithMountedDirectory(daggerInternalMkdocsWorkingDirectory, dir).
		WithWorkdir(daggerInternalMkdocsWorkingDirectory).
		WithEnvVariable("TINI_SUBREAPER", "true").
		WithEnvVariable("PIP_ROOT_USER_ACTION", "ignore")

	ctr = ctr.WithExec(append([]string{"python3", "-m", "pip", "install"}, packages...), dagger.ContainerWithExecOpts{
		SkipEntrypoint: true,
	})

	return ctr
}

// MkdocsService starts and returns an HTTP service mounted to directory for live reload and preview of mkdocs.
func (m *Garden) MkdocsService(ctx context.Context, dir *Directory) *Service {
	fmt.Printf("open locally at\n\t 👉 http://localhost:%d\n\t👉 http://127.0.0.1:%d\n\t👉 http://0.0.0.0:%d\n", localHostPort, localHostPort, localHostPort)
	return m.MkdocsBaseContainer(ctx, dir).
		WithWorkdir("/docs").
		WithMountedDirectory("/docs", dir).
		WithExec([]string{
			"mkdocs",
			"serve",
			"--clean",
			"--dev-addr",
			fmt.Sprintf("%s:%d", internalLocalHostAddressToUse, localHostPort),
		}, dagger.ContainerWithExecOpts{
			SkipEntrypoint: true,
		}).
		WithExposedPort(localHostPort, dagger.ContainerWithExposedPortOpts{
			Description: "mkdocs-local-server",
			Protocol:    dagger.Tcp,
		}).
		WithEnvVariable("TINI_SUBREAPER", "true").
		AsService()
}

func (m *Garden) MkdocsBuild(ctx context.Context, dir *Directory) *Directory {
	return m.MkdocsBaseContainer(ctx, dir).
		WithWorkdir("/docs").
		WithMountedDirectory("/docs", dir).
		WithExec([]string{"mkdocs", "build", "--clean", "--verbose", "--site-dir", qualifiedInternalArtifactsSiteDir}, dagger.ContainerWithExecOpts{
			SkipEntrypoint: true,
		}).Directory(qualifiedInternalArtifactsSiteDir)
}

// ======================================
// MARK: Mage To Do
// ======================================

// mage runs mage commands but inside dagger...
// this is from https://github.com/rancher/opni/blob/a4916c088ac0d1e94d56c2469e21b7b088a0be98/dagger/util.go#L44
// requires Go to be in image, and then mage was installed in it with this approach: https://github.com/rancher/opni/blob/a4916c088ac0d1e94d56c2469e21b7b088a0be98/dagger/main.go#L282C1-L289C38
// func mage(target []string, opts ...dagger.ContainerWithExecOpts) ([]string, dagger.ContainerWithExecOpts) {
// 	if len(opts) == 0 {
// 		opts = append(opts, dagger.ContainerWithExecOpts{})
// 	}
// 	mageCmd := []string{"mage", "-v"}
// 	return append(mageCmd, target...), opts[0]
// }

// // Serve starts the mkdocs service
// func (m *Garden) MkdocsServe(ctx context.Context, dir *Directory) (*Container, error) {
// 	s, err := m.MkdocsBaseContainer(ctx, dir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer s.Stop(ctx)

// 	return dag.Container().WithServiceBinding("mkdocs", s), nil
// }
