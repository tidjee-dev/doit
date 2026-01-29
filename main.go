package main

import (
	"os"

	cmd "github.com/tidjee-dev/doit/cmd/doit"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
