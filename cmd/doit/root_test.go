package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testTasksYAML = `app:
  name: doit
  version: v0.1.0
  description: test app
  main_file: main.go
  authors:
    - Tester
  repo_url: https://example.com/repo

tasks:
  hello:
    category: Test
    description: prints hello
    commands:
      - echo hello
`

func withCwd(t *testing.T, dir string) {
	t.Helper()

	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(old); err != nil {
			t.Fatalf("restore cwd: %v", err)
		}
	})
}

func writeTasksFile(t *testing.T, dir, content string) {
	t.Helper()

	path := filepath.Join(dir, "tasks.yml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write tasks.yml: %v", err)
	}
}

func runRoot(t *testing.T, dir string, args ...string) error {
	t.Helper()

	withCwd(t, dir)
	rootCmd.SetArgs(args)
	t.Cleanup(func() {
		rootCmd.SetArgs(nil)
	})

	return Execute()
}

func TestExecute_NoArgs_ShowsHelpPath(t *testing.T) {
	dir := t.TempDir()
	writeTasksFile(t, dir, testTasksYAML)

	if err := runRoot(t, dir); err != nil {
		t.Fatalf("execute without args should succeed, got: %v", err)
	}
}

func TestExecute_TaskRunsSuccessfully(t *testing.T) {
	dir := t.TempDir()
	writeTasksFile(t, dir, testTasksYAML)

	if err := runRoot(t, dir, "hello"); err != nil {
		t.Fatalf("execute task should succeed, got: %v", err)
	}
}

func TestExecute_TaskNotFound(t *testing.T) {
	dir := t.TempDir()
	writeTasksFile(t, dir, testTasksYAML)

	err := runRoot(t, dir, "missing")
	if err == nil || !strings.Contains(err.Error(), "task 'missing' not found") {
		t.Fatalf("expected missing task error, got: %v", err)
	}
}

func TestExecute_TooManyArgs(t *testing.T) {
	dir := t.TempDir()
	writeTasksFile(t, dir, testTasksYAML)

	err := runRoot(t, dir, "hello", "extra")
	if err == nil || !strings.Contains(err.Error(), "expected exactly one task name") {
		t.Fatalf("expected args validation error, got: %v", err)
	}
}
