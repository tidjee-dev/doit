package main

import (
	"os"

	"github.com/tidjee-dev/doit/internal/config"
	"github.com/tidjee-dev/doit/internal/runner"
	"github.com/tidjee-dev/doit/internal/ui"
)

func main() {
	cfg, err := config.Load("tasks.yml")
	if err != nil {
		ui.Error(err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		ui.PrintHelp(cfg)
		return
	}

	r := runner.New(cfg)

	if err := r.Run(os.Args[1]); err != nil {
		ui.Error(err.Error())
		os.Exit(1)
	}

	ui.PrintSummary(r.TaskCount(), r.Duration())
}
