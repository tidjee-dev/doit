---
title: Examples and Usage Patterns
sidebar_position: 1
---

This page shows practical `doit` usage patterns for real projects 🧪
All examples follow the explicit, single-dependency execution model.

All snippets below assume this required `app` block is present in `tasks.yml`:

```yaml
app:
  name: my-app
  version: v0.1.0
  description: My application
  main_file: main.go
  authors:
    - Your Name
  repo_url: https://github.com/yourname/yourrepo
```

## Basic Go Project

Typical Go build flow:

```yaml
env:
  BIN_DIR: bin

tasks:
  deps:
    category: Setup
    description: Download dependencies
    commands:
      - go mod download
      - go mod tidy

  build:
    category: Build
    description: Build binary
    depends_on:
      - deps
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }}

  run:
    category: Dev
    description: Run application
    depends_on:
      - deps
    commands:
      - go run {{ .App.MainFile }}
```

Execution chain:

```plain
deps → build
```

or

```plain
deps → run
```

Clear and readable 🔍

## Test Pipeline

Sequential validation pipeline:

```yaml
tasks:
  lint:
    category: Quality
    description: Run linter
    commands:
      - golangci-lint run

  test:
    category: Quality
    description: Run tests
    depends_on:
      - lint
    commands:
      - go test ./...

  coverage:
    category: Quality
    description: Coverage report
    depends_on:
      - test
    commands:
      - go test ./... -cover
```

Each stage blocks the next on failure ❌

## Build Variants with Task Env

Per-task overrides:

```yaml
env:
  BIN_DIR: bin

tasks:
  build-linux:
    category: Build
    description: Linux build
    env:
      GOOS: linux
      GOARCH: amd64
    commands:
      - go build -o {{ .Env.BIN_DIR }}/app-linux

  build-windows:
    category: Build
    description: Windows build
    env:
      GOOS: windows
      GOARCH: amd64
    commands:
      - go build -o {{ .Env.BIN_DIR }}/app.exe
```

Task env overrides global env 🎯

## Tooling Wrapper Pattern

Wrap external tools consistently:

```yaml
tasks:
  migrate-up:
    category: Database
    description: Apply migrations
    commands:
      - migrate -path db/migrations up

  migrate-down:
    category: Database
    description: Rollback last migration
    commands:
      - migrate -path db/migrations down 1
```

Provides a stable interface for team usage 👥

## Explicit Release Chain

Readable release pipeline:

```yaml
tasks:
  deps:
    category: Release
    description: Prepare deps
    commands:
      - go mod download

  build:
    category: Release
    description: Build release
    depends_on:
      - deps
    commands:
      - go build -ldflags "-s -w"

  package:
    category: Release
    description: Create archive
    depends_on:
      - build
    commands:
      - tar -czf release.tar.gz bin/
```

No hidden graph, no implicit ordering 📦

## Anti-Pattern — Multi Dependency Simulation

Avoid:

```yaml
depends_on:
  - deps
  - lint
```

Not allowed by design 🚫

Instead chain:

```plain
deps → lint → build
```

## Pattern Guidelines

Good patterns:

- Linear chains
- Small focused tasks
- Clear naming
- Explicit ordering

Avoid:

- Wide dependency trees
- Hidden side effects
- Overloaded tasks
- Long command strings
