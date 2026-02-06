---
title: Sprig Helper Reference
sidebar_position: 2
---

This page lists the Sprig helpers that are **recommended** when using templates in `doit`, along with their intended use cases.

Sprig provides many functions, but only a subset makes sense in a task runner context.
This table focuses on helpers that improve **readability, safety, and maintainability**.

## String Helpers

| Helper            | Purpose                       | Typical use in `doit`         |
| ----------------- | ----------------------------- | ----------------------------- |
| `upper` / `lower` | Change string case            | Normalize flags, tags, output |
| `title`           | Title-case a string           | Human-readable output         |
| `trim`            | Remove surrounding whitespace | Clean env values              |
| `trimPrefix`      | Remove a prefix               | Strip `v` from versions       |
| `trimSuffix`      | Remove a suffix               | Strip extensions              |
| `replace`         | Replace substrings            | File names, env vars          |
| `substr`          | Extract substring             | Short IDs, version slicing    |
| `repeat`          | Repeat a string               | Output formatting             |

### Sorting

| Helper | Purpose                | Typical use in `doit` |
| ------ | ---------------------- | --------------------- |
| `sort` | Sort a list of strings | CLI args, display     |

## Defaulting & Presence

These helpers are **core** to safe template usage.

| Helper     | Purpose               | Typical use in `doit` |
| ---------- | --------------------- | --------------------- |
| `default`  | Fallback value        | Optional env vars     |
| `empty`    | Test zero values      | Presence checks       |
| `coalesce` | First non-empty value | Fallback chains       |

> Prefer `default` and `empty` over `if` whenever possible.

## List Helpers

| Helper  | Purpose        | Typical use in `doit` |
| ------- | -------------- | --------------------- |
| `list`  | Create a slice | Build arguments       |
| `join`  | Join elements  | CLI args, display     |
| `first` | First element  | Primary value         |
| `last`  | Last element   | Fallback value        |
| `len`   | Length of list | Sanity checks         |

## Dict / Map Helpers

Use sparingly, but useful for structured lookups.

| Helper   | Purpose               | Typical use in `doit` |
| -------- | --------------------- | --------------------- |
| `dict`   | Create a map          | Inline lookup tables  |
| `get`    | Get value by key      | Optional config       |
| `hasKey` | Key existence         | Feature flags         |
| `keys`   | List keys             | Debug / reporting     |
| `values` | List values           | Debug / reporting     |
| `pluck`  | Extract key from maps | Structured data ⚠️    |

> ⚠️ `pluck` expects **multiple maps**, not a list.

## Flow & Logic Helpers

Use for **simple conditions only**.

| Helper      | Purpose           | Typical use in `doit` |
| ----------- | ----------------- | --------------------- |
| `ternary`   | Conditional value | Flags, toggles        |
| `eq` / `ne` | Equality checks   | Env switching         |
| `gt` / `lt` | Comparisons       | Limits, guards        |

> `ternary` requires a **boolean condition**.
> Non-empty strings are truthy in Go templates. Compare string flags explicitly when needed.

## Math Helpers

Keep usage minimal.

| Helper        | Purpose    | Typical use in `doit` |
| ------------- | ---------- | --------------------- |
| `add` / `sub` | Arithmetic | Counters              |
| `mul` / `div` | Arithmetic | Simple math           |
| `max` / `min` | Bounds     | Limits                |

If math becomes complex, move logic to a script.

## Time & Formatting

| Helper   | Purpose       | Typical use in `doit` |
| -------- | ------------- | --------------------- |
| `now`    | Current time  | Logging               |
| `date`   | Format time   | Artifact names        |
| `printf` | Format output | Aligned logs, demos   |

`printf` is strongly recommended for readable output.

## Recommendation Summary

- ✅ Formatting helpers: **encouraged**
- ✅ Defaulting helpers: **essential**
- ⚠️ Dict & flow helpers: **use carefully**
- ❌ Complex logic: **avoid**
