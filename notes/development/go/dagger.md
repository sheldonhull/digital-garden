---
title: Dagger
slug: dagger
lastmod: 2023-03-07 19:48
date: 2023-03-07 19:20
tags:
  - containers
  - go
  - build-release-engineering
---

[Containerized magic with Go and BuildKit](https://dagger.io?ref=sheldonhull.com)

Will put some experience notes here soon as I've successfully built angular and nginx containers with it and overall was a great experience.
With the upcoming services support, I can see a whole lot more usage cases too.

## Example Building An Angular Project

Using mage, here's an example of how to invoke Mage to build an angular project without any angular tooling installed locally.

```go
const AngularVersion = "15"

// Build runs the angular build via Dagger.
func (Dagger) Build(ctx context.Context) error {
  client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
  if err != nil {
    pterm.Error.Printfln("unable to connect to dagger: %s", err)
    return err
  }
  defer client.Close()

  homedir, err := os.UserHomeDir()
  if err != nil {
    return err
  }
  npm := client.Container().From("node:lts-alpine")
  npm = npm.WithMountedDirectory("/src", client.Host().Directory(".")).
    WithWorkdir("/src")

  path := "dist/"
  npm = npm.WithExec([]string{"npm", "install", "-g", fmt.Sprintf("@angular/cli@%s", AngularVersion)})
  npm = npm.WithExec([]string{"ng", "config", "-g", "cli.warnings.versionMismatch", "false"})
  npm = npm.WithExec([]string{"ng", "v"})
  npm = npm.WithExec([]string{"npm", "ci"})
  npm = npm.WithExec([]string{"ng", "build", "--configuration", "production"})

  // Copy "dist/" from container to host.
  _, err = npm.Directory(path).Export(ctx, path)
  if err != nil {
    return err
  }
  return nil
}
```

??? example "example of handling both local and CI private npm auth"

    Here you can handle both running in a CI context or a remote context by evaluting for a CI variable that would provide back a CI system generated `.npmrc`.
    If this isn't provided, mount the file from the home directory into the build container.

    Note this container isn't for publishing, it's a build container copying the `dist/` contents back to the project directory.

    ```go
    npmrcFile := &dagger.Secret{}

    // bypassing any mounting of npmrc, as CI tooling should update any private  inline with current file here
    if os.Getenv("NPM_CONFIG_USERCONFIG") != "" {
      pterm.Info.Printfln("[OVERRIDE] NPM_CONFIG_USERCONFIG: %s", os.Getenv("NPM_CONFIG_USERCONFIG"))
      npmrcDir := filepath.Dir(os.Getenv("NPM_CONFIG_USERCONFIG"))
    } else {
      // [DEFAULT] NPM config set from home/.npmrc
      npmrcFile = client.Host().Directory(homedir, dagger.HostDirectoryOpts{Include: []string{".npmrc"}}).File(".npmrc").Secret()

      // if npmrcFile doesn't exist output error
      if _, err := os.Stat(filepath.Join(homedir, ".npmrc")); os.IsNotExist(err) {
        return errors.New("missing npmrc file")
      }
      npm = npm.WithMountedSecret("/root/.npmrc", npmrcFile)
    }
    ```
