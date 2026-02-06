---
title: Templates
sidebar_position: 1
---

`doit` renders all command strings using Go’s `text/template` engine **before execution**.

This allows commands to be dynamically generated based on application metadata, environment variables, and task context.

## Template Context

The following objects are available inside templates:

### `.App`

Metadata defined in the `app` section.

```gotemplate
{{ .App.Name }}
{{ .App.Version }}
```

### `.Env`

Merged environment variables, resolved in this order (last wins):

1. Process environment
2. Global `env`
3. Task-level `env`

```gotemplate
{{ .Env.PATH }}
{{ .Env.BIN_DIR }}
```

### `.Task`

Metadata of the currently executed task:

- `Name`
- `Category`
- `Description`

```gotemplate
{{ .Task.Name }}
{{ .Task.Category }}
```

## Missing Keys & Zero Values

Template rendering is **safe by default**.

Missing keys do **not** cause errors and resolve to Go zero values:

| Type   | Zero value |
| ------ | ---------- |
| string | `""`       |
| bool   | `false`    |
| number | `0`        |

This enables optional lookups without guards:

```yaml
commands:
  - echo '{{ default "normal" .Env.MISSING }}'
```

⚠️ **Important:**
In Go templates, non-empty strings are truthy. If you want boolean semantics for string flags, compare explicitly (for example `eq .Env.DEBUG "true"`).

## Enabling Sprig Helpers

By default, only standard `text/template` functions are available.

To enable [Sprig](https://masterminds.github.io/sprig/) helpers, opt in explicitly:

```yaml
templates:
  sprig: true
```

Once enabled, all Sprig **text/template** helpers are available in command strings.

## Quick Sprig Examples

```yaml
commands:
  - echo '{{ "doit" | upper }}'
  - echo '{{ default "bin" .Env.BIN_DIR }}'
  - echo '{{ list "a" "b" | join ", " }}'
  - echo '{{ add 2 3 }}'
```

## References

- Go `text/template`: [https://pkg.go.dev/text/template](https://pkg.go.dev/text/template)
- Sprig Documentation: [https://masterminds.github.io/sprig/](https://masterminds.github.io/sprig/)
