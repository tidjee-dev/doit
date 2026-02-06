---
title: CLI Reference
sidebar_position: 1
---

This page documents the `doit` command-line interface 🖥️

The CLI is intentionally small and explicit. Commands map directly to task execution and configuration setup.

## Base Command

```bash
doit
```

Running without arguments lists all available tasks grouped by category.

Output includes:

- Category name
- Task name
- Task description

This is the primary discovery view 📋

## Run a Task

```bash
doit <task_name>
```

Example:

```bash
doit build
```

Behavior:

- Validates `tasks.yml`
- Resolves dependency (if defined)
- Expands templates (Go templates; optional Sprig helpers)
- Executes commands sequentially
- Stops on first failure

## Init Command

```bash
doit init
```

Creates a base `tasks.yml` in the current directory.

Behavior:

- Writes a starter configuration
- Does not overwrite an existing file
- Shows a warning if file already exists

Use this to bootstrap new projects ⚙️

## Execution Flow

When running a task:

1. Load configuration file
2. Validate configuration rules
3. Resolve dependency (max one)
4. Merge environment variables
5. Resolve templates
6. Execute commands in order

No hidden steps occur.

## Exit Codes

`doit` returns standard process exit codes:

| Situation        | Exit Code |
| ---------------- | --------- |
| Task success     | 0         |
| Validation error | non-zero  |
| Command failure  | non-zero  |
| Missing task     | non-zero  |

This allows safe CI usage 🧪

## Output Behavior

CLI output is formatted for readability:

- Task headers
- Command status indicators
- Success markers
- Failure markers

Command stdout and stderr are streamed directly.

## Working Directory

Tasks execute in:

```plain
current working directory
```

There is no directory switching unless your command does it explicitly.

## CI Usage

Typical CI usage:

```bash
doit build
doit test
```

Because execution is deterministic and validation-first, failures are immediate and visible.
