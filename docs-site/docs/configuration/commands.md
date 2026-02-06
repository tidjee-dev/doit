---
title: Commands
sidebar_position: 7
---

The `commands` field defines the shell commands executed by a task 🖥️

It is required for every task and runs sequentially.

## Structure

```yaml
tasks:
  build:
    category: Build
    description: Compile app
    commands:
      - go build
      - echo done
```

## Execution Model

Commands are executed:

- In the order defined
- Sequentially
- In the system shell
- With merged environment variables
- After dependency execution (if defined)

Execution stops immediately on failure ❌

## Failure Behavior

If any command exits with a non-zero status:

- Task execution stops
- Remaining commands are skipped
- doit returns a non-zero exit code

There is no partial success state.

## Shell Behavior

Commands are passed to the system shell:

- Linux / macOS → `sh -c`
- Windows → `cmd /C`

Avoid shell-specific features when portability matters ⚠️

## Templating Support

Commands support template expressions.

Core context:

```plain
{{ .App.* }}
{{ .Env.* }}
{{ .Task.* }}
```

Example:

```yaml
commands:
  - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }}
```

Templates are resolved before command execution.
For missing-key behavior and Sprig helpers, see `Configuration > Templates and Sprig`.

## Output

Command output is streamed to the console:

- stdout shown
- stderr shown
- command success/failure indicated

Formatting is handled by the CLI UI layer.

## Best Practices

- Keep commands explicit ✍️
- Avoid hidden side effects
- Prefer multiple simple commands over one complex chain
- Do not rely on shell aliases
