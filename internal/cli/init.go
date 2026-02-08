package cli

import (
	"fmt"
	"os"

	"github.com/tidjee-dev/doit/internal/ui"
)

const tasksFile = "tasks.yml"

func Init() error {
	_, err := os.Stat(tasksFile)
	if err == nil {
		ui.Warn(fmt.Sprintf("%s already exists at current directory", tasksFile))
		return nil
	}

	if !os.IsNotExist(err) {
		return fmt.Errorf("stat %s failed: %w", tasksFile, err)
	}

	if err := os.WriteFile(tasksFile, []byte(defaultTasksYAML), 0644); err != nil {
		return fmt.Errorf("write %s failed: %w", tasksFile, err)
	}

	ui.Success(fmt.Sprintf("created %s at current directory", tasksFile))
	return nil
}

const defaultTasksYAML = `# yaml-language-server: $schema=https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
# See docs at https://tidjee-dev.github.io/doit

app:
  name: my-app
  version: v0.1.0
  description: My application
  main_file: main.go
  authors:
    - Your Name
  repo_url: https://github.com/yourname/yourrepo

env:
  BIN_DIR: bin

tasks:
  deps:
    category: Dependencies
    description: Install dependencies
    commands:
      - go mod tidy

  build:
    category: Build
    description: Compile the application
    depends_on:
      - deps
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} {{ .App.MainFile }}

  dev:
    category: Development
    description: Run the application
    depends_on:
      - deps
    quiet: true
    commands:
      - go run {{ .App.MainFile }}
`
