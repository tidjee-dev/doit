---
title: doit
sidebar_position: 1
---

[fr]

**doit** is a simple, fast, and explicit task runner written in Go ⚡

It lets you define project tasks using YAML with strict validation and predictable execution.

It is designed as a minimal and opinionated alternative to Makefile — focused on clarity, structure, and determinism.

## Why doit

Most task runners grow complex over time.

doit intentionally stays small and strict:

- Explicit execution order
- Strict schema validation
- No implicit behavior
- One dependency maximum per task
- No hidden execution graph

If a task chain is not obvious by reading the file, it is considered wrong.

## Goals

- Explicit execution order
- Strict configuration validation
- Minimal feature surface
- Predictable runtime behavior
- Easy to reason about execution flow

## Non-Goals

- No dynamic task graphs
- No wildcard dependencies
- No shell-specific tricks
- No implicit chaining
- No magic resolution rules

## Core Features

- Single Go binary 🧩
- YAML configuration (`tasks.yml`)
- Explicit task dependency
- Task categories for clean listing
- Global and task-level environment variables
- JSON Schema validation
- Template support for commands
- Clean CLI output
- Zero runtime dependencies

## Configuration Overview

doit uses a single configuration file:

```plain
tasks.yml
```

Main sections:

- `app` — metadata and template context
- `env` — global environment variables
- `tasks` — executable task definitions

Each task defines commands and optional dependency — execution is always sequential and explicit.
