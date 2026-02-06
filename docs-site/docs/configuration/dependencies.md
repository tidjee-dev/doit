---
title: Dependencies
sidebar_position: 6
---

Tasks can declare a dependency using the `depends_on` field 🔗

A dependency is executed before the task itself.

## Structure

```yaml
tasks:
  build:
    category: Build
    description: Compile app
    depends_on:
      - deps
    commands:
      - go build
```

## Rules

- `depends_on` is optional
- Accepts an array with **maximum one item**
- The referenced task must exist
- Self-dependency is invalid
- Cycles are rejected at runtime

Valid:

```yaml
depends_on:
  - deps
```

Invalid:

```yaml
depends_on:
  - deps
  - lint
```

## Execution Order

If a dependency is defined:

```plain
dependency → task
```

Example:

```plain
deps → build → run
```

Each task runs only after its dependency completes successfully.

## No Dependency Graph Resolution

doit does not build a DAG and does not resolve multi-branch graphs.

There is:

- No implicit ordering
- No dependency merging
- No parallel execution
- No wildcard dependency matching

This is intentional for predictability 🧭

## Chaining Pattern

Create readable chains by linking tasks one by one:

```yaml
tasks:
  deps: ...

  build:
    depends_on:
      - deps

  test:
    depends_on:
      - build
```

This keeps execution flow obvious from top to bottom.

## Anti-Patterns

Avoid:

- Deep hidden chains
- Reusing one task as dependency everywhere
- Simulating DAG behavior

If ordering becomes unclear, split tasks differently.
