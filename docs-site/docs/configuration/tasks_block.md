---
title: Tasks Block
sidebar_position: 4
---

The `tasks` block defines all executable tasks 🧩

Each task contains metadata, commands, and optional configuration. Tasks are executed by name from the CLI.

## Structure

```yaml
tasks:
  task_name:
    category: Build
    description: Compile the project
    commands:
      - go build
```

The key under `tasks` is the task name used in the CLI.

Run with:

```bash
doit task_name
```

## Task Name Rules

Task names must:

- Start with a letter
- Contain only: `a-z A-Z 0-9 _ -`
- Be unique within the file

Valid:

```yaml
build
build-api
test_all
```

Invalid:

```yaml
1build
build*
```

## Required Fields

Each task must define:

- `category`
- `description`
- `commands`

Example:

```yaml
tasks:
  deps:
    category: Dependencies
    description: Install dependencies
    commands:
      - go mod tidy
```

## Optional Fields

A task may also define:

- `depends_on` — single dependency
- `env` — task-level environment variables
- `quiet` — hide doit logs for that task

Example:

```yaml
tasks:
  build:
    category: Build
    description: Compile binary
    depends_on:
      - deps
    env:
      CGO_ENABLED: "0"
    commands:
      - go build
```

## Command Execution

Commands are executed:

- In declared order
- Sequentially
- In the system shell
- With merged environment variables

Execution stops on the first failing command ❌

## CLI Listing

Tasks are displayed grouped by category when running:

```bash
doit
```

Categories are labels only — they do not affect execution order.

## Validation Rules

- Missing required fields cause failure
- Empty command lists are rejected
- Unknown fields are rejected at runtime
- Dependency limits are enforced

Validation occurs before execution starts.
