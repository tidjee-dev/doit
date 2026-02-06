package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const validYAML = `app:
  name: doit
  version: v0.1.0
  description: A task runner written in Go
  main_file: main.go
  authors:
    - Tidjee
  repo_url: https://github.com/tidjee-dev/doit

env:
  BIN_DIR: bin

tasks:
  build:
    category: Build
    description: Compile the application
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}
`

func writeConfig(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.yml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	return path
}

func TestLoad_ValidConfig(t *testing.T) {
	path := writeConfig(t, validYAML)
	if _, err := Load(path); err != nil {
		t.Fatalf("expected valid config, got error: %v", err)
	}
}

func TestLoad_MissingAppFields(t *testing.T) {
	path := writeConfig(t, strings.ReplaceAll(validYAML, "name: doit\n", "name: \n"))
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "app.name is required") {
		t.Fatalf("expected app.name error, got: %v", err)
	}
}

func TestLoad_InvalidVersion(t *testing.T) {
	bad := strings.ReplaceAll(validYAML, "version: v0.1.0", "version: 0.1")
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "app.version must be semver") {
		t.Fatalf("expected semver error, got: %v", err)
	}
}

func TestLoad_InvalidMainFile(t *testing.T) {
	bad := strings.ReplaceAll(validYAML, "main_file: main.go", "main_file: main.txt")
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "app.main_file must be a .go path") {
		t.Fatalf("expected main_file error, got: %v", err)
	}
}

func TestLoad_InvalidRepoURL(t *testing.T) {
	bad := strings.ReplaceAll(validYAML, "repo_url: https://github.com/tidjee-dev/doit", "repo_url: not-a-url")
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "app.repo_url must be a valid URL") {
		t.Fatalf("expected repo_url error, got: %v", err)
	}
}

func TestLoad_InvalidEnvKey(t *testing.T) {
	bad := strings.ReplaceAll(validYAML, "BIN_DIR: bin", "1INVALID: bin")
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "env key") {
		t.Fatalf("expected env key error, got: %v", err)
	}
}

func TestLoad_InvalidTaskName(t *testing.T) {
	bad := strings.ReplaceAll(validYAML, "build:", "1build:")
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "task name") {
		t.Fatalf("expected task name error, got: %v", err)
	}
}

func TestLoad_TaskWithNoCommands(t *testing.T) {
	bad := strings.ReplaceAll(validYAML,
		"commands:\n      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}\n",
		"commands: []\n",
	)
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "commands must not be empty") {
		t.Fatalf("expected commands error, got: %v", err)
	}
}

func TestLoad_DependsOnMissingTask(t *testing.T) {
	bad := strings.ReplaceAll(validYAML,
		"commands:\n      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}\n",
		"depends_on:\n      - missing\n    commands:\n      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}\n",
	)
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "depends_on") {
		t.Fatalf("expected depends_on missing error, got: %v", err)
	}
}

func TestLoad_DependsOnTooMany(t *testing.T) {
	bad := strings.ReplaceAll(validYAML,
		"commands:\n      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}\n",
		"depends_on:\n      - deps\n      - other\n    commands:\n      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}\n",
	)
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "depends_on must have at most one entry") {
		t.Fatalf("expected depends_on max error, got: %v", err)
	}
}

func TestLoad_NoTasks(t *testing.T) {
	emptyTasks := `app:
  name: doit
  version: v0.1.0
  description: A task runner written in Go
  main_file: main.go
  authors:
    - Tidjee
  repo_url: https://github.com/tidjee-dev/doit

env:
  BIN_DIR: bin

tasks: {}
`
	path := writeConfig(t, emptyTasks)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "no tasks defined") {
		t.Fatalf("expected no tasks error, got: %v", err)
	}
}

func TestLoad_UnknownTopLevelField(t *testing.T) {
	bad := validYAML + "\nunknown_field: true\n"
	path := writeConfig(t, bad)
	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "field unknown_field not found") {
		t.Fatalf("expected unknown field error, got: %v", err)
	}
}
