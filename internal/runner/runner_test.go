package runner

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/tidjee-dev/doit/internal/config"
)

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	os.Exit(0)
}

func newTestCommander(t *testing.T, calls *[]string) func(string) *exec.Cmd {
	t.Helper()
	return func(cmd string) *exec.Cmd {
		*calls = append(*calls, cmd)
		c := exec.Command(os.Args[0], "-test.run=TestHelperProcess", "--")
		c.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
		return c
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = old

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("read stdout: %v", err)
	}
	_ = r.Close()

	return string(out)
}

func TestRun_ExecutesDependenciesInOrder(t *testing.T) {
	cfg := config.Config{
		App: config.AppConfig{
			Name:        "doit",
			Version:     "v0.1.0",
			Description: "test",
			MainFile:    "main.go",
			Authors:     []string{"Tester"},
			RepoURL:     "https://example.com/repo",
		},
		Tasks: map[string]config.Task{
			"deps": {
				Category:    "Setup",
				Description: "deps",
				Commands:    []string{"echo {{ .Task.Name }}"},
			},
			"build": {
				Category:    "Build",
				Description: "build",
				DependsOn:   []string{"deps"},
				Commands:    []string{"echo {{ .Task.Name }}"},
			},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	if err := r.Run("build", Verbose); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(calls) != 2 {
		t.Fatalf("expected 2 commands, got %d: %#v", len(calls), calls)
	}
	if calls[0] != "echo deps" || calls[1] != "echo build" {
		t.Fatalf("unexpected order: %#v", calls)
	}
}

func TestRun_CycleDetected(t *testing.T) {
	cfg := config.Config{
		Tasks: map[string]config.Task{
			"a": {Category: "Test", Description: "a", DependsOn: []string{"b"}, Commands: []string{"noop"}},
			"b": {Category: "Test", Description: "b", DependsOn: []string{"a"}, Commands: []string{"noop"}},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	err := r.Run("a", Verbose)
	if err == nil || !strings.Contains(err.Error(), "cyclic dependency detected") {
		t.Fatalf("expected cycle error, got: %v", err)
	}
}

func TestRun_TaskNotFound(t *testing.T) {
	cfg := config.Config{Tasks: map[string]config.Task{}}
	r := New(cfg)
	err := r.Run("missing", Verbose)
	if err == nil || !strings.Contains(err.Error(), "task 'missing' not found") {
		t.Fatalf("expected not found error, got: %v", err)
	}
}

func TestRun_MissingTemplateKeyRendersZeroValue(t *testing.T) {
	cfg := config.Config{
		Tasks: map[string]config.Task{
			"echo-missing": {
				Category:    "Test",
				Description: "missing template key",
				Commands:    []string{"echo {{ .Env.MISSING }}"},
			},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	if err := r.Run("echo-missing", Verbose); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(calls) != 1 {
		t.Fatalf("expected 1 command, got %d: %#v", len(calls), calls)
	}
	if calls[0] != "echo " {
		t.Fatalf("expected rendered zero value for missing key, got: %q", calls[0])
	}
}

func TestRun_QuietTaskSuppressesDependencyLogs(t *testing.T) {
	cfg := config.Config{
		Tasks: map[string]config.Task{
			"deps": {
				Category:    "Setup",
				Description: "deps",
				Commands:    []string{"echo deps"},
			},
			"build": {
				Category:    "Build",
				Description: "build",
				DependsOn:   []string{"deps"},
				Quiet:       true,
				Commands:    []string{"echo build"},
			},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	out := captureStdout(t, func() {
		if err := r.Run("build", Verbose); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if strings.Contains(out, "echo deps") || strings.Contains(out, "echo build") || strings.Contains(out, "╭") {
		t.Fatalf("expected no doit task logs in quiet mode, got output: %q", out)
	}
}

func TestRun_SprigFunctionsEnabledWhenConfigured(t *testing.T) {
	cfg := config.Config{
		Templates: config.TemplatesConfig{Sprig: true},
		Tasks: map[string]config.Task{
			"trim": {
				Category:    "Test",
				Description: "sprig",
				Commands:    []string{`echo {{ trim "  hi  " }}`},
			},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	if err := r.Run("trim", Verbose); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(calls) != 1 || calls[0] != "echo hi" {
		t.Fatalf("expected sprig-trimmed command, got: %#v", calls)
	}
}

func TestRun_SprigFunctionsDisabledByDefault(t *testing.T) {
	cfg := config.Config{
		Tasks: map[string]config.Task{
			"trim": {
				Category:    "Test",
				Description: "sprig",
				Commands:    []string{`echo {{ trim "  hi  " }}`},
			},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	err := r.Run("trim", Verbose)
	if err == nil || !strings.Contains(err.Error(), `function "trim" not defined`) {
		t.Fatalf("expected missing sprig function error, got: %v", err)
	}
}

func TestRun_TemplateIncludesTaskDescription(t *testing.T) {
	cfg := config.Config{
		Tasks: map[string]config.Task{
			"echo-desc": {
				Category:    "Test",
				Description: "my description",
				Commands:    []string{`echo {{ .Task.Description }}`},
			},
		},
	}

	r := New(cfg)
	var calls []string
	r.commander = newTestCommander(t, &calls)

	if err := r.Run("echo-desc", Verbose); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(calls) != 1 || calls[0] != "echo my description" {
		t.Fatalf("expected task description in template context, got: %#v", calls)
	}
}
