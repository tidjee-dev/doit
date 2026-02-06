---
title: Installation
sidebar_position: 3
---

This page covers all supported installation methods for `doit` 📦

## Requirements

- Go 1.25.6 or newer recommended
- Git (for source installs)

Check:

```bash
go version
```

## Supported Platforms

| Platform | `go install` | From Source |
| -------- | :----------: | :---------: |
| Linux    |      ✅      |     ✅      |
| macOS    |      ✅      |     ✅      |
| Windows  |      ✅      |     ✅      |

## Install with `go install` (Recommended)

```bash
go install github.com/tidjee-dev/doit@latest
```

This builds and installs the binary into:

```plain
$GOPATH/bin
```

or

```plain
$HOME/go/bin
```

Ensure this directory is in your `PATH`.

Verify installation:

```bash
doit
```

### If '@latest' is not yet available ⏳

Right after a new release tag, the Go module proxy may take a few minutes to update.

If needed, install directly from Git:

```bash
GOPROXY=direct go install github.com/tidjee-dev/doit@latest
```

## Install from Source

Clone the repository:

```bash
git clone https://github.com/tidjee-dev/doit.git
cd doit
```

Build using Go:

```bash
go build -o bin/doit
```

Run directly:

```bash
./bin/doit
```

### Linux | macOS Global Install

Move the binary into a system path:

```bash
sudo mv bin/doit /usr/local/bin/doit
```

Verify:

```bash
doit
```

### Windows Global Install

Build:

```bash
go build -o bin/doit.exe
```

Add the binary directory to your system `PATH`.

Then run:

```bash
doit
```

## Upgrade

Using Go install:

```bash
go install github.com/tidjee-dev/doit@latest
```

This overwrites the existing binary with the newest version.

## Install specific version

Using Go install:

```bash
go install github.com/tidjee-dev/doit@v0.1.0
```
