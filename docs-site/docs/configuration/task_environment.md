---
title: Task Environment
sidebar_position: 5
---

Tasks can define their own environment variables using a task-level `env` block 🌿

These variables apply only to that task and override global environment values when keys overlap.

## Structure

```yaml
tasks:
  build:
    category: Build
    description: Compile binary
    env:
      CGO_ENABLED: "0"
      OUTPUT_DIR: dist
    commands:
      - go build -o {{ .Env.OUTPUT_DIR }}/app
```

## Scope

Task environment variables are:

- Available only to that task
- Available to templates
- Injected into the command process environment
- Merged with global env values

## Merge Order

Environment resolution order:

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

Result inside the task:

```plain
BIN_DIR = dist
```

## Type Rules

Same rules as global env.

## Template Access

Task env variables are accessed through:

```plain
{{ .Env.KEY }}
```

Example:

```yaml
commands:
  - echo {{ .Env.OUTPUT_DIR }}
```

## Best Practices

- Use task env for per-task tuning 🎯
- Avoid redefining many globals
- Keep names consistent across tasks
- Do not store secrets in tasks.yml 🔒
