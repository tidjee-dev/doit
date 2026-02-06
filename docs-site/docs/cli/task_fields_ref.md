---
title: Task Fields Reference
sidebar_position: 2
---

This page provides a field-by-field reference for task definitions 📘

Each task is defined under the `tasks` block and keyed by its task name.

## Task Definition Structure

```yaml
tasks:
  build:
    category: Build
    description: Compile app
    depends_on:
      - deps
    env:
      CGO_ENABLED: "0"
    commands:
      - go build
```

## `category`

- Type: string
- Required: yes
- Description: logical grouping label
- Used only for CLI display

Example:

```yaml
category: Build
```

Tasks are grouped by category when listing tasks.

No execution behavior is tied to category.

## `description`

- Type: string
- Required: yes
- Description: short task summary shown in CLI

Example:

```yaml
description: Compile the application
```

Keep descriptions short and action-oriented ✍️

## `depends_on`

- Type: string array
- Required: no
- Max items: **1**
- Description: single task dependency

Example:

```yaml
depends_on:
  - deps
```

Rules:

- Referenced task must exist
- Only one dependency allowed
- Cycles are rejected at runtime
- Runs before the task

## `env`

- Type: object (string → string)
- Required: no
- Description: task-level environment variables

Example:

```yaml
env:
  CGO_ENABLED: "0"
  OUTPUT_DIR: dist
```

Rules:

- Keys must be valid identifiers
- Values must be strings
- Overrides global env on conflict
- Available in templates

Template usage:

```plain
{{ .Env.OUTPUT_DIR }}
```

## `quiet`

- Type: boolean
- Required: no
- Default: `false`
- Description: hide doit task logs while still streaming command output

Example:

```yaml
quiet: true
```

## `commands`

- Type: string array
- Required: yes
- Min items: 1
- Description: commands executed sequentially

Example:

```yaml
commands:
  - go build
  - echo done
```

Rules:

- Order is preserved
- Stops on first failure
- Executed in system shell
- Supports templating

## Validation Constraints

Task definitions are rejected if:

- Required fields are missing ❌
- commands list is empty ❌
- depends_on has more than one item ❌
- quiet is not a boolean ❌
- Field types are incorrect ❌
- Unknown fields are rejected at runtime ❌

Validation occurs before execution.

## Recommended Pattern

Keep each task:

- Focused 🎯
- Small
- Explicit
- Readable without cross-referencing many tasks
