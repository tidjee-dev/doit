package runner

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"github.com/tidjee-dev/doit/internal/config"
)

type TemplateContext struct {
	App  config.AppConfig
	Env  map[string]string
	Task struct {
		Name        string
		Category    string
		Description string
	}
}

func renderTemplate(
	raw string,
	cfg config.Config,
	taskName string,
	task config.Task,
) (string, error) {

	mergedEnv := mergeEnv(cfg.Env, task.Env)

	mergedEnvMap := make(map[string]string, len(mergedEnv))
	for _, e := range mergedEnv {
		key, value := splitEnv(e)
		mergedEnvMap[key] = value
	}

	ctx := TemplateContext{
		App: cfg.App,
		Env: mergedEnvMap,
	}

	ctx.Task.Name = taskName
	ctx.Task.Category = task.Category
	ctx.Task.Description = task.Description

	tpl := template.New("command").
		Option("missingkey=zero")

	if cfg.Templates.Sprig {
		tpl = tpl.Funcs(sprig.TxtFuncMap())
	}

	tpl, err := tpl.Parse(raw)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, ctx); err != nil {
		return "", err
	}

	return buf.String(), nil
}
