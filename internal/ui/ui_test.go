package ui

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/tidjee-dev/doit/internal/config"
)

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

func TestSortedKeys(t *testing.T) {
	got := sortedKeys(map[string][]string{
		"Build": {"build"},
		"Test":  {"test"},
		"Deps":  {"deps"},
	})

	want := []string{"Build", "Deps", "Test"}
	if len(got) != len(want) {
		t.Fatalf("unexpected key count: got %d want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("unexpected order at %d: got %q want %q", i, got[i], want[i])
		}
	}
}

func TestPrintHelp_IncludesUsageAndTasks(t *testing.T) {
	cfg := config.Config{
		App: config.AppConfig{Name: "doit"},
		Tasks: map[string]config.Task{
			"deps": {
				Category:    "Dependencies",
				Description: "Install dependencies",
			},
			"build-z": {
				Category:    "Build",
				Description: "Build binary",
			},
			"build-a": {
				Category:    "Build",
				Description: "Build alpha",
			},
		},
	}

	out := captureStdout(t, func() {
		PrintHelp(cfg)
	})

	for _, expected := range []string{
		"doit",
		"Task runner for project: doit",
		"Usage",
		"doit <task>",
		"Available tasks",
		"Dependencies",
		"Build",
		"deps",
		"build-z",
		"build-a",
		"Install dependencies",
		"Build binary",
		"Build alpha",
	} {
		if !strings.Contains(out, expected) {
			t.Fatalf("expected output to contain %q, got:\n%s", expected, out)
		}
	}

	if strings.Index(out, "build-a") > strings.Index(out, "build-z") {
		t.Fatalf("expected deterministic alphabetical order within category, got:\n%s", out)
	}
}

func TestPrintTaskAndSummaryOutput(t *testing.T) {
	out := captureStdout(t, func() {
		PrintTaskHeader("Build", "build", "Compile app")
		PrintCommand("go test ./...")
		PrintTaskFooter(1500 * time.Millisecond)
		PrintSummary(2, 1500*time.Millisecond)
	})

	for _, expected := range []string{
		"Build",
		"build",
		"Compile app",
		"go test ./...",
		"Completed in",
		"2 tasks executed in",
	} {
		if !strings.Contains(out, expected) {
			t.Fatalf("expected output to contain %q, got:\n%s", expected, out)
		}
	}
}
