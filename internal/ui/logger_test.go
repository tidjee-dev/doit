package ui

import (
	"strings"
	"testing"
)

func TestLoggerHelpers_PrintExpectedPrefixes(t *testing.T) {
	out := captureStdout(t, func() {
		Success("created")
		Warn("existing")
		Info("details")
		Error("failed")
	})

	for _, expected := range []string{
		"[SUCCESS]",
		"[WARN]",
		"[INFO]",
		"[ERROR]",
		"created",
		"existing",
		"details",
		"failed",
	} {
		if !strings.Contains(out, expected) {
			t.Fatalf("expected output to contain %q, got:\n%s", expected, out)
		}
	}
}
