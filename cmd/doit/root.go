package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/tidjee-dev/doit/internal/config"
	"github.com/tidjee-dev/doit/internal/runner"
	"github.com/tidjee-dev/doit/internal/ui"
)

var rootCmd = &cobra.Command{
	Use:           "doit <task>",
	Short:         "Doit is a simple task runner",
	SilenceUsage:  true,
	SilenceErrors: false,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("expected exactly one task name")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cfg, err := config.Load("tasks.yml")
			if err != nil {
				return err
			}
			ui.PrintHelp(cfg)
			return nil
		}

		cfg, err := config.Load("tasks.yml")
		if err != nil {
			return err
		}

		r := runner.New(cfg)

		if err := r.Run(args[0]); err != nil {
			return err
		}

		ui.PrintSummary(r.TaskCount(), r.Duration())
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
