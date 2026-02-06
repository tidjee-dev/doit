---
title: Configuration Overview
sidebar_position: 1

---

`doit` uses a single configuration file named **`tasks.yml`** ⚙️

This file defines application metadata, environment variables, and executable tasks.
It is validated before execution, and the published JSON Schema keeps editor feedback aligned with runtime rules.

## File Location

The file must exist in the working directory where `doit` is executed:

```plain
tasks.yml
```

Initialize it with:

```bash
doit init
```

## Top-Level Structure

The configuration has three top-level sections:

```yaml
app:
templates:
env:
tasks:
```

* `app` — required — metadata and template context
* `templates` — optional — template engine settings
* `env` — optional — global environment variables
* `tasks` — required — task definitions

For template behavior and Sprig setup, see `Configuration > Templates and Sprig`.

## Example Structure

```yaml
app:
  name: my-app
  version: v0.1.0
  description: Example project
  main_file: main.go
  authors:
    - Your Name
  repo_url: https://github.com/your/repo

env:
  BIN_DIR: bin

tasks:
  build:
    category: Build
    description: Compile the application
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}
```

## Validation

Runtime validation is performed by `doit` before running any task.
JSON Schema is provided for editor tooling and early feedback while editing.

Benefits:

* Early error detection
* Field validation
* Type checking
* Editor autocomplete support 🧭

Schema URL:

```plain
https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
```

## Execution Rules

* Unknown fields are rejected at runtime
* Missing required fields cause failure
* Invalid types cause failure
* Dependency limits are enforced

Execution never starts if validation fails.
