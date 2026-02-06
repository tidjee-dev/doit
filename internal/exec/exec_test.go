package exec

import (
	"runtime"
	"testing"
)

func TestCommand_UsesExpectedShell(t *testing.T) {
	cmd := Command("echo ok")

	if runtime.GOOS == "windows" {
		if cmd.Args[0] != "cmd" || cmd.Args[1] != "/C" {
			t.Fatalf("expected cmd /C on windows, got args: %#v", cmd.Args)
		}
		return
	}

	if cmd.Args[0] != "sh" || cmd.Args[1] != "-c" {
		t.Fatalf("expected sh -c on non-windows, got args: %#v", cmd.Args)
	}
}
