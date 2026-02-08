# doit

Explicit task runner written in Go projects.

- [Docs](https://tidjee-dev.github.io/doit)
- [Schema](https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json)

## Prerequisites

- Go 1.25.6 or newer

## Install

```bash
go install github.com/tidjee-dev/doit@latest
```

## Quick Start

```bash
doit init
doit
doit build
```

## tasks.yml

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
app:
  name: my-app
  version: v0.1.0
  description: My application
  main_file: main.go
  authors:
    - Your Name
  repo_url: https://github.com/yourname/yourrepo

templates:
  sprig: true

env:
  BIN_DIR: bin

tasks:
  deps:
    category: Dependencies
    description: Install dependencies
    commands:
      - go mod tidy

  build:
    category: Build
    description: Build binary
    depends_on:
      - deps
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}
```

## Notes

- `depends_on` supports at most one dependency.
- `quiet: true` hides doit logs only (command output is still streamed).
- Templates use Go `text/template` with `.App`, `.Env`, `.Task` context.
- Missing template keys resolve to zero values.
- Set `templates.sprig: true` to enable full Sprig helpers.

## Development

```bash
go test ./...
```

When changing behavior, add or update `*_test.go` files in the affected package.
