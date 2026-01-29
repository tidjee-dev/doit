# doit

A **simple, fast, and explicit task runner** written in Go.

`doit` lets you define project tasks in YAML with **clear structure**, **strict validation**, and **predictable execution**.
It is designed as a **minimal and opinionated alternative to Makefile**.

## Why `doit`

Most task runners grow complex over time.

`doit` deliberately stays small:

* ✅ Explicit execution order
* ✅ Strict schema validation
* ✅ No implicit magic
* ✅ One dependency max
* ❌ No hidden task graphs
* ❌ No shell-specific tricks

If a task chain is not obvious by reading the file, it’s probably wrong.

## Features

* ⚡ Single Go binary
* 📄 YAML configuration (`tasks.yml`)
* 🔗 Explicit task dependency (`depends_on`)
* 🧩 Task categories for clean listing
* 🌱 Global and task-level environment variables
* 🧪 JSON Schema validation
* 🖥️ Clean, readable CLI output
* 🛠️ Zero runtime dependencies

## Installation

### Using `go install` (recommended)

```bash
go install github.com/tidjee-dev/doit@latest
```

Make sure `$GOPATH/bin` (or `$HOME/go/bin`) is in your `PATH`.

### From source

```bash
git clone https://github.com/tidjee-dev/doit.git
cd doit
go run main.go build
```

Optional global install (Linux / macOS):

```bash
sudo mv doit /usr/local/bin/doit
```

Windows: add the binary directory to your `PATH`.

## Usage

Initialize a new project:

```bash
doit init
```

> Creates a base `tasks.yml` in the current directory (if it does not already exist).

Run a task:

```bash
doit <task>
```

List all tasks:

```bash
doit
```

## Configuration

`doit` uses a single configuration file: **`tasks.yml`**

It is strictly validated using JSON Schema.

Schema URL:

```text
https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
```

## File Structure

```yaml
app:     # required
env:     # optional (global env)
tasks:   # required
```

## `app`

Application metadata and entry point.

```yaml
app:
  name: doit
  version: v0.1.0
  description: A task runner written in Go
  main_file: main.go
  author: Donatien Pinet
  repo_url: https://github.com/tidjee-dev/doit
```

### `app` Fields

| Field         | Type     | Description       |
| ------------- | -------- | ----------------- |
| `name`        | string   | Application name  |
| `version`     | string   | SemVer (`vX.Y.Z`) |
| `description` | string   | Short description |
| `main_file`   | string   | Go entry file     |
| `author`      | string[] | Maintainer        |
| `repo_url`    | string   | Repository URL    |

## `env` (optional)

Global environment variables injected into **all tasks**.

```yaml
env:
  BIN_DIR: bin
```

Rules:

* Keys must be valid shell identifiers
* Values are strings only

## `tasks`

Tasks are indexed by name and executed via the CLI.

```yaml
tasks:
  deps:
    category: Dependencies
    description: Install and tidy dependencies
    commands:
      - go mod download
      - go mod tidy
```

### Task Name Rules

* Must start with a letter
* Allowed characters: `a-z A-Z 0-9 _ -`

## Task Definition

```yaml
tasks:
  build:
    category: Build
    description: Compile the application
    depends_on:
      - deps
    env:
      CGO_ENABLED: "0"
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}
```

### `task` Fields

| Field         | Required  | Type     | Notes                 |
| ------------- | :------:  | -------- | --------------------- |
| `category`    |     ✅    | string   | Logical grouping      |
| `description` |     ✅    | string   | Displayed in CLI      |
| `depends_on`  |     ❌    | string[] | **Max 1 dependency**  |
| `env`         |     ❌    | object   | Task-local env vars   |
| `commands`    |     ✅    | string[] | Executed sequentially |

## Templating

`doit` supports simple templating inside commands.

Available contexts:

* `{{ .App.* }}` → values from `app`
* `{{ .Env.* }}` → merged env (global + task)

Example:

```yaml
go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}
```

## Dependency Model

* `depends_on` accepts **only one task**
* Dependency runs **before** the task
* No implicit chaining
* No DAG resolution

Example:

```yaml
depends_on:
  - deps
```

This constraint is **intentional** to keep execution predictable.

## Validation of `tasks.yml`

Benefits:

* Autocomplete
* Inline errors
* Early feedback before runtime

## Example `tasks.yml`

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
# See docs at https://github.com/tidjee-dev/doit

app:
  name: my-app
  version: v0.1.0
  description: My application
  main_file: main.go
  author: Your Name
  repo_url: https://github.com/yourname/yourrepo

env:
  BIN_DIR: bin

tasks:
  deps:
    category: Dependencies
    description: Install and tidy dependencies
    commands:
      - go mod download
      - go mod tidy

  build:
    category: Build
    description: Compile the application for the current platform
    depends_on:
      - deps
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}

  dev:
    category: Development
    description: Run the application
    depends_on:
      - deps
    commands:
      - go run {{ .App.MainFile }}
```

## Design Principles

* ❌ No implicit behavior
* ❌ No wildcard dependencies
* ❌ No dynamic task graphs
* ✅ Explicit execution order
* ✅ Strict configuration
* ✅ Easy to reason about

`doit` favors **clarity over flexibility**.

## License

MIT
