# Contributing

Thanks for contributing to **doit**. The goal is to keep the project **simple, explicit, and predictable**.

## Scope

Contributions should:

- Improve correctness, clarity, or DX
- Reduce ambiguity or hidden behavior
- Keep the core small (features over configuration are discouraged)

Please avoid:

- Magic defaults or implicit behavior
- Non-deterministic execution
- Features that duplicate shell or Make behavior

## Development Setup

Requirements:

- Go (latest stable)

Clone and run tests:

```bash
go test ./...
```

## Code Style

- Follow standard Go conventions (`gofmt`, idiomatic naming)
- Prefer small, explicit types over generic helpers
- Avoid reflection unless strictly necessary
- Errors should be explicit and actionable

## Tasks & Execution

- Task resolution must remain deterministic
- Validation should fail fast with clear error messages
- Any change to task behavior must be documented

## Schema Changes

If you modify the `tasks.yml` structure, please update `schema.json` accordingly:

- Keep backward compatibility when possible
- Update documentation accordingly
- Add validation tests covering the change

## Commits

- Keep commits focused and minimal
- One logical change per commit
- Write clear commit messages (imperative, present tense)

## Pull Requests

Before opening a PR:

- Ensure `go test ./...` passes
- Update docs if behavior changes
- Explain _why_ the change is needed, not just _what_ changed

If unsure about a feature or direction, open an issue first.
