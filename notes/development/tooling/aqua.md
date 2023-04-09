---
title: aqua
description: A cli version manager for the discriminating cli connoisseur.
date: 2023-04-04 12:56
tags:
  - cli
  - tooling
categories: []
lastmod: 2023-04-04 13:10
---

## Aqua Overview

A cli version manager for the discriminating cli connoisseur, this tool is great to install binaries both at global level and project level.

## Quick Start

[Quick Start](https://aquaproj.github.io/docs/tutorial) includes install commands to setup.

## Update Your Path

=== "macOS/Linux"

      ```shell title="$HOME/.zshenv"
      export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
      ```

=== "windows"

      ```powershell title="Windows: System Environment Variables (Outside Process Scope)"
      [Environment]::SetEnvironmentVariable('PATH', ((Join-Path $ENV:LOCALAPPDATA 'aquaproj-aqua' 'bin') , $ENV:PATH -join [IO.Path]::PathSeparator), 'Machine')
      ```

      ```powershell title="$PROFILE.CurrentUserAllHosts"
      $ENV:PATH = ((Join-Path $ENV:LOCALAPPDATA 'aquaproj-aqua' 'bin') , $ENV:PATH -join [IO.Path]::PathSeparator)
      ```

## Global Tooling Setup

To create these files, navigate to the directory and run `aqua init && aqua init-policy`.

```shell title="$HOME/.zshenv"
export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
export AQUA_GLOBAL_CONFIG=${XDG_CONFIG_HOME:-$HOME/.config}/aqua/aqua.yaml
export AQUA_POLICY_CONFIG=${XDG_CONFIG_HOME:-$HOME/.config}/aqua/aqua-policy.yaml
```

## For Windows