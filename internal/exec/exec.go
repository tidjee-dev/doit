package exec

import (
	"os/exec"
	"runtime"
)

func Command(cmd string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", cmd)
	}
	return exec.Command("bash", "-c", cmd)
}
