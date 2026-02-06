---
title: Schema Validation
sidebar_position: 9
---

`tasks.yml` is validated before any task runs 📄

Runtime validation is performed by `doit`. JSON Schema is published for editor support and static checks.

## Schema URL

Official schema:

[https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json](https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json)

This schema defines:

- Required fields
- Allowed field types
- Naming rules
- Dependency limits
- Structure constraints

## Enable Editor Validation

Add this header at the top of `tasks.yml`:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
```

Supported by editors using YAML Language Server:

- VS Code
- Neovim (yaml-ls)
- JetBrains IDEs
- Other LSP-compatible editors

## Validation Coverage

Runtime and schema rules cover:

- Top-level structure
- Required blocks (`app`, `tasks`)
- Field types
- Unknown field detection
- Task name format
- Command arrays
- Environment value types
- `depends_on` max length = 1

## Example — Valid

```yaml
tasks:
  build:
    category: Build
    description: Compile
    commands:
      - go build
```

## Example — Invalid

Missing required field:

```yaml
tasks:
  build:
    commands:
      - go build
```

Too many dependencies:

```yaml
depends_on:
  - deps
  - lint
```

Wrong type:

```yaml
env:
  PORT: 8080
```

Must be:

```yaml
env:
  PORT: "8080"
```

## Runtime Behavior

Validation occurs before execution:

- File is parsed
- Runtime rules are applied
- Errors are reported
- Execution stops if invalid ❌

No partial execution occurs on invalid configuration.

## Best Practices

- Always enable schema in editor 🧭
- Treat validation errors as build errors
- Keep schema URL pinned to the repo version if needed
- Validate before committing changes
