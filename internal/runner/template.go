package runner

import (
	"bytes"
	"text/template"

	"github.com/tidjee-dev/doit/internal/config"
)

type TemplateContext struct {
	App  config.AppConfig
	Env  map[string]string
	Task struct {
		Name     string
		Category string
	}
}

func renderTemplate(
	raw string,
	cfg config.Config,
	taskName string,
	task config.Task,
) (string, error) {

	ctx := TemplateContext{
		App: cfg.App,
		Env: cfg.Env,
	}
	ctx.Task.Name = taskName
	ctx.Task.Category = task.Category

	tpl, err := template.New("command").Parse(raw)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}
