---
title: Sprig Pitfalls & Common Mistakes
sidebar_position: 3
---

Sprig helpers are powerful, but Go `text/template` has **strict typing rules** and explicit truthiness rules.

This page documents the **most common mistakes** when using Sprig in `doit`, why they happen, and the correct patterns to use instead.

## 1. Boolean-Like Strings Can Be Misleading

⚠️ **Misleading**

```gotemplate
{{ if .Env.DEBUG }}debug{{ end }}
```

If `DEBUG` is a non-empty string like `"false"`, this still renders `debug`.

✅ **Correct**

```gotemplate
{{ if eq .Env.DEBUG "true" }}debug{{ end }}
```

or

```gotemplate
{{ if not (empty .Env.DEBUG) }}debug{{ end }}
```

**Rule:**
Non-empty strings are truthy. Compare string flags explicitly or use `empty` for presence checks.

## 2. `ternary` Requires a Boolean

❌ **Wrong**

```gotemplate
{{ ternary "set" "unset" .Env.MISSING }}
```

Error:

```text
expected bool; got string
```

✅ **Correct**

```gotemplate
{{ ternary "unset" "set" (empty .Env.MISSING) }}
```

or

```gotemplate
{{ ternary "set" "unset" (ne .Env.MISSING "") }}
```

**Rule:**
`ternary`’s condition must evaluate to a `bool`.

## 3. `pluck` Does Not Accept Lists

❌ **Wrong**

```gotemplate
{{ pluck "name" (list (dict "name" "a") (dict "name" "b")) }}
```

Error:

```text
expected map[string]interface {}; got []interface {}
```

✅ **Correct**

```gotemplate
{{ pluck "name" (dict "name" "a") (dict "name" "b") }}
```

**Rule:**
`pluck` expects **multiple maps as arguments**, not a slice.

## 4. `coalesce` Can Return `<nil>`

❌ **Surprising**

```gotemplate
{{ coalesce .Env.A .Env.B }}
```

Output:

```text
<no value>
```

This happens when **all values are empty**.

✅ **Safe pattern**

```gotemplate
{{ default "fallback" (coalesce .Env.A .Env.B) }}
```

**Rule:**
Always provide a final fallback when using `coalesce`.

## 5. Zero Values Are Silent

Missing keys resolve to zero values:

| Type   | Zero value |
| ------ | ---------- |
| string | `""`       |
| bool   | `false`    |
| number | `0`        |

```gotemplate
{{ .Env.MISSING }}   → ""
```

This is safe, but **typos are silent**.

✅ **Debug tip**

```gotemplate
{{ printf "%q" .Env.MISSING }}
```

## 6. Avoid Complex Logic in Templates

❌ **Bad idea**

```gotemplate
{{ if and (gt (len .Env.PATH) 10) (eq .Task.Name "build") }}
```

✅ **Better**

Move logic into a script and keep templates declarative:

```yaml
commands:
  - ./scripts/build.sh {{ .Env.PATH }}
```

## Summary Rules

- ❗ Non-empty strings are truthy
- ❗ `ternary` needs a boolean
- ❗ `pluck` ≠ list
- ⚠️ `coalesce` may return `<nil>`
- ✅ Prefer `default` and `empty`
- ❌ Avoid complex logic

These rules will prevent **90% of Sprig-related errors**.
