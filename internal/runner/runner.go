package runner

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/tidjee-dev/doit/internal/config"
	exec_util "github.com/tidjee-dev/doit/internal/exec"
	"github.com/tidjee-dev/doit/internal/ui"
)

type ExecMode uint8

const (
	Quiet ExecMode = iota
	Verbose
)

type Runner struct {
	cfg       config.Config
	visited   map[string]bool
	running   map[string]bool
	startTime time.Time
	taskCount int

	commander func(string) *exec.Cmd
}

func New(cfg config.Config) *Runner {
	return &Runner{
		cfg:       cfg,
		visited:   make(map[string]bool),
		running:   make(map[string]bool),
		startTime: time.Now(),
		commander: exec_util.Command,
	}
}

func (r *Runner) Run(taskName string, parentMode ExecMode) error {
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
	defer delete(r.running, taskName)

	mode := parentMode
	if task.Quiet {
		mode = Quiet
	}

	for _, dep := range task.DependsOn {
		if err := r.Run(dep, mode); err != nil {
			return err
		}
	}

	if err := r.execute(taskName, task, mode); err != nil {
		return err
	}

	r.visited[taskName] = true
	r.taskCount++

	return nil
}

func (r *Runner) execute(name string, task config.Task, mode ExecMode) error {
	start := time.Now()

	if mode == Verbose {
		ui.PrintTaskHeader(
			task.Category,
			name,
			task.Description,
		)
	}

	env := mergeEnv(r.cfg.Env, task.Env)

	for _, rawCmd := range task.Commands {
		cmdStr, err := renderTemplate(rawCmd, r.cfg, name, task)
		if err != nil {
			return fmt.Errorf("template error in task '%s': %w", name, err)
		}

		if mode == Verbose {
			ui.PrintCommand(cmdStr)
		}

		cmd := r.commander(cmdStr)
		cmd.Env = env

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("task '%s' failed: %w", name, err)
		}
	}

	if mode == Verbose {
		ui.PrintTaskFooter(time.Since(start))
	}

	return nil
}

func (r *Runner) TaskCount() int {
	return r.taskCount
}

func (r *Runner) Duration() time.Duration {
	return time.Since(r.startTime)
}
