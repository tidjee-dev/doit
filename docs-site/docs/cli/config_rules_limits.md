---
title: Configuration Rules and Limits
sidebar_position: 3
---

`doit` enforces strict configuration rules to keep execution predictable and easy to reason about 🧭

These limits are intentional design choices, not missing features.

## Global Constraints

The configuration file must:

- Be named `tasks.yml`
- Exist in the working directory
- Contain required top-level blocks
- Pass runtime validation before execution

Unknown top-level fields are rejected at runtime ❌

## Required Top-Level Blocks

These blocks must exist:

- `app`
- `tasks`

Optional:

- `env`

Missing required blocks cause validation failure.

## Task Constraints

Each task must include:

- `category`
- `description`
- `commands` (minimum one command)

A task may include:

- `depends_on` (max one)
- `env`

## Dependency Limit

Dependencies are limited to **one per task** 🔗

Rules:

- No multiple dependencies
- No dependency arrays > 1
- No automatic graph resolution
- No parallel dependency execution

Chaining must be explicit.

## Command Constraints

Commands must be:

- A non-empty string array
- Executable in the system shell
- Valid after template resolution

Empty command lists are rejected.

## Environment Constraints

Environment blocks (`env`) must:

- Be key/value maps
- Use string values only
- Use valid identifier keys

Invalid:

```yaml
env:
  PORT: 8080
```

Valid:

```yaml
env:
  PORT: "8080"
```

## Template Constraints

- Templates use Go `text/template`
- Context includes `.App.*`, `.Env.*`, `.Task.*`
- Missing keys resolve to zero values
- Sprig helpers are optional via `templates.sprig: true`

Detailed reference:

- `Configuration > Templates and Sprig`

## Naming Constraints

Task names must:

- Start with a letter
- Use only: letters, numbers, `_`, `-`
- Be unique

Invalid names are rejected at validation time.

## Execution Guarantees

Given a valid configuration:

- Execution order is explicit
- No hidden task execution
- No implicit dependencies
- No runtime graph building
- No partial execution on validation error

Behavior is stable and predictable.

## Recommended Practices

- Prefer explicit chains over shared dependencies
- Keep tasks small and focused
- Avoid deep dependency chains
- Use schema validation in your editor 🧩
