---
title: App Block
sidebar_position: 2
---

The `app` block defines application metadata and provides values available to templates 📘

It is **required** and must appear at the top level of `tasks.yml`.

## Structure

```yaml
app:
  name: my-app
  version: v0.1.0
  description: My application
  main_file: main.go
  authors:
    - Your Name
  repo_url: https://github.com/your/repo
```

## Fields

### `name`

- Type: string
- Required: yes
- Description: Application name
- Min length: 1

Example:

```yaml
name: doit
```

Template usage:

```plain
{{ .App.Name }}
```

### `version`

- Type: string
- Required: yes
- Description: Application version
- Expected format: SemVer with prefix (`vX.Y.Z`)

Example:

```yaml
version: v1.2.0
```

### `description`

- Type: string
- Required: yes
- Description: Short description of the application
- Min length: 1

Example:

```yaml
description: My application
```

### `main_file`

- Type: string
- Required: yes
- Description: Path to the Go entry file
- Have to be a go file

Example:

```yaml
main_file: main.go
```

Template usage:

```plain
{{ .App.MainFile }}
```

### `authors`

- Type: array
- Required: yes
- Description: Project author or maintainer
- Min items: 1

#### `authors` Item

- Type: string
- Required: yes
- Description: Author or maintainer name
- Min length: 1

Example:

```yaml
authors:
  - Your Name
```

### `repo_url`

- Type: string
- Required: yes
- Description: Source repository URL
- Expected format: URI

Example:

```yaml
repo_url: https://github.com/your/repo
```

## Template Context

All `app` fields are available inside command templates:

Examples:

```plain
{{ .App.Name }}
{{ .App.Version }}
{{ .App.MainFile }}
```

These values are resolved before command execution.

## Validation Rules

- All fields are required
- String fields must be non-empty, and `authors` must contain at least one item
- Unknown fields are rejected at runtime
- Invalid URLs fail validation
