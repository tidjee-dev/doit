package cli

import (
	"fmt"
	"os"

	"github.com/tidjee-dev/doit/internal/ui"
)

const tasksFile = "tasks.yml"

func Init() error {
	if _, err := os.Stat(tasksFile); err == nil {
		ui.Error(fmt.Sprintf("%s already exists at current directory", tasksFile))
		return nil
	} else {
		os.WriteFile(tasksFile, []byte(defaultTasksYAML), 0644)
		ui.Success(fmt.Sprintf("created %s at current directory", tasksFile))
		return nil
	}
}

const defaultTasksYAML = `# yaml-language-server: $schema=https://raw.githubusercontent.com/tidjee-dev/doit/main/schema.json
# See docs at https://github.com/tidjee-dev/doit

app:
  name: my-app
  version: v0.1.0
  description: My application
  main_file: main.go
  author:
		- Your Name
  repo_url: https://github.com/yourname/yourrepo

env:
  BIN_DIR: bin

tasks:
  deps:
    category: Dependencies
    description: Install and tidy dependencies
    commands:
      - go mod download
      - go mod tidy

  build:
    category: Build
    description: Compile the application
    depends_on:
      - deps
    commands:
      - go build -o {{ .Env.BIN_DIR }}/{{ .App.Name }} main.go

  dev:
    category: Development
    description: Run the application
    depends_on:
      - deps
    commands:
      - go run main.go
`
