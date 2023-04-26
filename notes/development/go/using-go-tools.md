---
title: Using Go Tools
date: 2023-04-19 18:12
lastmod: 2023-04-19 18:12
---

This is primarily focused on folks who don't use Go tooling everday, but want to use Go tools.
Maybe you need help getting up and running?

## Go Binaries

Tools that can compile to a Go binary such as CLI tools or a web server can be installed from source easily, by running `go install`.

There's a few things you need though to do this.

- Go installed 😀
  - Using aqua makes this easy.
- The path the binaries are dropped into isn't in your `PATH` by default, so you need to ensure your shell of choice has this path added so the binaries can be found globally.
  - `export PATH="$(go env GOPATH)/bin:${PATH}"
- The right way to invoke it.