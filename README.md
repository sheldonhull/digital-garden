# Digital Garden

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sheldonhull/digital-garden/mkdocs-publish.yml?style=for-the-badge)

An experiment on moving my "digital garden" content from hugo to mkdocs.
While my primary blog is still on hugo, the effort to maintain a personal knowledge base on it is too high.

Instead, I'm trying out mkdocs which I've used in the past for several projects and will see how if it reduces the friction to write.

## Run

- `aqua install`
- `mage serve`

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
