package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func withWorkingDir(t *testing.T, dir string) {
	t.Helper()

	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir to temp dir: %v", err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(old); err != nil {
			t.Fatalf("restore cwd: %v", err)
		}
	})
}

func TestInit_CreatesDefaultTasksYAML(t *testing.T) {
	dir := t.TempDir()
	withWorkingDir(t, dir)

	if err := Init(); err != nil {
		t.Fatalf("init failed: %v", err)
	}

	path := filepath.Join(dir, tasksFile)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read created file: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "https://tidjee-dev.github.io/doit") {
		t.Fatalf("expected docs website link in default config, got:\n%s", content)
	}
	if !strings.Contains(content, "authors:") {
		t.Fatalf("expected authors field in default config, got:\n%s", content)
	}
	if !strings.Contains(content, "quiet: true") {
		t.Fatalf("expected quiet mode in default config, got:\n%s", content)
	}
	if !strings.Contains(content, "sprig: true") {
		t.Fatalf("expected sprig templates enabled in default config, got:\n%s", content)
	}
	if !strings.Contains(content, `echo "hello {{ .App.Authors | first }}"`) {
		t.Fatalf("expected hello demo task in default config, got:\n%s", content)
	}
	if !strings.Contains(content, "go run {{ .App.MainFile }}") {
		t.Fatalf("expected MainFile template in default config, got:\n%s", content)
	}
}

func TestInit_DoesNotOverwriteExistingTasksYAML(t *testing.T) {
	dir := t.TempDir()
	withWorkingDir(t, dir)

	const existing = "tasks: {}\n"
	if err := os.WriteFile(tasksFile, []byte(existing), 0644); err != nil {
		t.Fatalf("seed tasks.yml: %v", err)
	}

	if err := Init(); err != nil {
		t.Fatalf("init should not fail when file exists: %v", err)
	}

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		t.Fatalf("read tasks.yml: %v", err)
	}

	if string(data) != existing {
		t.Fatalf("expected existing file to stay unchanged, got:\n%s", string(data))
	}
}
