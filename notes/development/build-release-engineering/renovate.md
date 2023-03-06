---
title: renovate
description: Use renovate for dependency management
date: 2023-03-06 17:18
tags:
  - dependency-management
slug: renovate
---

## Run Locally Via Docker

### For Azure DevOps

!!! note "Git Safe Directory"

    This doesn't mount to the host `/tmp`, unlike examples in GitHub, because it flags the directory as owned by the docker user (likely root), which causes Git's safe directory feature to block.

    Easier to just disable mounting to the host in this scenario, as `git config --global --add safe.directory /tmp/renovate/repos/*` didn't seem to work.

Set the environment variables: `export AZURE_DEVOPS_ORG=foo`, and the other `AZURE_DEVOPS_EXT_PAT`, and finally replace `PROJECTNAME/REPO`.

```shell
docker run --rm -it \
    -e RENOVATE_PLATFORM="azure" \
    -e RENOVATE_ENDPOINT="https://dev.azure.com/${AZURE_DEVOPS_ORG}/" \
    -e GITHUB_COM_TOKEN=$(gh auth token) \
    -e SYSTEM_ACCESSTOKEN=$AZURE_DEVOPS_EXT_PAT \
    -e RENOVATE_TOKEN=$AZURE_DEVOPS_EXT_PAT \
    -e RENOVATE_DRY_RUN=full \
    -e LOG_LEVEL=debug \
    -v ${PWD}/config.js:/usr/src/app/config.js \
    -v /var/run/docker.sock:/var/run/docker.sock \
    renovate/renovate:latest --include-forks=false --dry-run=full PROJECTNAME/REPO

```
