# Digital Garden

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sheldonhull/digital-garden/mkdocs-publish.yml?style=for-the-badge)

An experiment on moving my "digital garden" content from hugo to mkdocs.
While my primary blog is still on hugo, the effort to maintain a personal knowledge base on it is too high.

Instead, I'm trying out mkdocs which I've used in the past for several projects and will see how if it reduces the friction to write.

## Dagger

- build: `dagger call mkdocs-build --dir ${PWD} export --path .artifacts`
- server: `dagger call mkdocs-service --dir . up --ports 8000:8000`

## Devcontainer

- Open in devcontainer or codespaces.
- Open `zsh-login` terminal, not bash, to ensure all the environment paths are set.
- Run `direnv allow` to load default environment variables.
- Run `mage init` to setup anything missing.

> If anything like Go/Mage are missing, run `aqua install --tags first && aqua install` to automatically fix that.

## Mage Commands

Get going by running `mage job:up` to run the docker pull, build, and run the local serve command.
It will ouput the url, which defaults to [locahost](http://localhost:3001)

```text
Targets:
  docker:build            🔨 Run docker build.
  docker:pull             pulls the squidfunk/mkdocs-material:latest Docker image
  docker:stop             🛑 Stop the mkdocs dockerized container.
  init                    ✔️ Init sets up the local tooling for writing and building.
  job:up                  can get everything running from scratch and server locally.
  mkdocs:build            Run mkdocs commands contained in docker
  mkdocs:ghDeploy         Run mkdocs commands contained in docker
  mkdocs:serve            🌐 Run mkdocs serve via Docker.
  trunk:init              ⚙️ Init installs trunk and ensures the plugins are setup.
  trunk:install           ⚙️ InstallTrunk installs trunk.io tooling if it isn't already found.
  trunk:installPlugins    ⚙️ InstallPlugins ensures the required runtimes are installed.
  trunk:upgrade           ⚙️ Upgrade upgrades trunk using itself and also the plugins.
```

## Why

- I want simple markdown.
- Ability to iterate quickly on small notes.
- No third part build required so I can be more picky on my publishing triggers (such as netlify which is great, but too busy).
- No need to manage directory structure and navigation/TOC builds.
- Potentionally replacing [docs](https://www.sheldonhull/docs?ref=digital-garden-repo).

## Remaining

- [ ] Backlinks possible?
- [ ] RSS for updates to mailbrew
- [ ] Gist styling isn't taking what I did on my personal blog, what does this need to render correctly?
- [ ] Search optimization with Algolia or leave it as is?
- [ ] Flat urls even though organized with hiearchy?