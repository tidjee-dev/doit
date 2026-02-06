---
title: Global Environment
sidebar_position: 3
---

The `env` block defines global environment variables injected into all tasks 🌱

It is optional and defined at the top level of `tasks.yml`.

## Structure

```yaml
env:
  BIN_DIR: bin
  CGO_ENABLED: "0"
```

All variables defined here are available to every task.

## Scope

Global environment variables are:

- Available to all commands
- Available to templates
- Merged with task-level environment variables
- Passed to the command process environment

## Type Rules

### Keys

- Keys must be valid shell identifiers
  - Allowed characters: `[a-zA-Z0-9_]`
- Keys must be unique

### Values

- Values must be strings
- Numbers and booleans must be quoted

Valid:

```yaml
env:
  PORT: "8080"
  DEBUG: "true"
```

Invalid:

```yaml
env:
  PORT: 8080
  DEBUG: true
```

## Template Usage

Global env variables are available via:

```plain
{{ .Env.KEY }}
```

Example:

```yaml
commands:
  - go build -o {{ .Env.BIN_DIR }}/app
```

## Merge Behavior

When a task defines its own `env`, the merge order is:

```plain
global env → task env
```

Task-level values override global ones.

Example:

```yaml
env:
  BIN_DIR: bin

tasks:
  build:
    env:
      BIN_DIR: dist
```

Result inside task:

```plain
BIN_DIR = dist
```

## Best Practices

- Keep global env small
- Use for shared constants
- Avoid secrets in tasks.yml 🔒
