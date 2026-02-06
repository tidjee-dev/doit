---
title: Quick Start
sidebar_position: 2
---

This guide gets you from zero to a working `doit` task in a few minutes 🚀

## Install

Recommended method:

```bash
go install github.com/tidjee-dev/doit@latest
```

Ensure your Go bin directory is in your `PATH`:

```bash
$HOME/go/bin
```

Verify:

```bash
doit --help
```

## Initialize a Project

Create a base tasks file in the current directory:

```bash
doit init
```

This creates:

```plain
tasks.yml
```

If the file already exists, initialization is skipped to avoid overwrite.

## Minimal Tasks File

Example:

```yaml
app:
  name: demo
  version: v0.1.0
  description: Demo app
  main_file: main.go
  authors:
    - Your Name
  repo_url: https://github.com/yourname/yourrepo

tasks:
  hello:
    category: Demo
    description: Print hello
    commands:
      - echo hello
```

## List Tasks

Run without arguments:

```bash
doit
```

Output shows tasks grouped by category 📋

## Run a Task

```bash
doit hello
```

Commands execute sequentially in the order defined.

## Add a Dependency

```yaml
tasks:
  deps:
    category: Setup
    description: Download modules
    commands:
      - go mod download

  run:
    category: Dev
    description: Run app
    depends_on:
      - deps
    commands:
      - go run main.go
```

Execution order:

```plain
deps → run
```

Only one dependency is allowed per task by design.
