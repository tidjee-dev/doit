package runner

import (
	"fmt"
	"os"
	"time"

	"github.com/tidjee-dev/doit/internal/config"
	exec_util "github.com/tidjee-dev/doit/internal/exec"
	"github.com/tidjee-dev/doit/internal/ui"
)

type ExecMode int

const (
	Verbose ExecMode = iota
	Quiet
)

type Runner struct {
	cfg       config.Config
	visited   map[string]bool
	running   map[string]bool
	startTime time.Time
	taskCount int

	execMode ExecMode
}

func New(cfg config.Config) *Runner {
	return &Runner{
		cfg:       cfg,
		visited:   make(map[string]bool),
		running:   make(map[string]bool),
		startTime: time.Now(),
	}
}

func (r *Runner) Run(taskName string) error {
	task, ok := r.cfg.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%s' not found", taskName)
	}

	if r.visited[taskName] {
		return nil
	}

	if r.running[taskName] {
		return fmt.Errorf("cyclic dependency detected at '%s'", taskName)
	}

	r.running[taskName] = true

	for _, dep := range task.DependsOn {
		if err := r.Run(dep); err != nil {
			return err
		}
	}

	if err := r.execute(taskName, task); err != nil {
		return err
	}

	r.running[taskName] = false
	r.visited[taskName] = true
	r.taskCount++

	return nil
}

func (r *Runner) execute(name string, task config.Task) error {
	start := time.Now()

	ui.PrintTaskHeader(
		task.Category,
		name,
		task.Description,
	)

	for _, cmdStr := range task.Commands {
		cmd := exec_util.Command(cmdStr)
		if r.execMode == Quiet {
			cmd.Stdout = nil
			cmd.Stderr = nil
		} else {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("task '%s' failed: %w", name, err)
		}

		ui.PrintCommand(cmdStr)
	}

	ui.PrintTaskFooter(time.Since(start))
	return nil
}

func (r *Runner) TaskCount() int {
	return r.taskCount
}

func (r *Runner) Duration() time.Duration {
	return time.Since(r.startTime)
}
