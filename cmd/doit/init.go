package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tidjee-dev/doit/internal/cli"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a base tasks.yml in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.Init(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
