package config

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	MainFile    string   `yaml:"main_file"`
	Authors     []string `yaml:"authors"`
	RepoURL     string   `yaml:"repo_url"`
}

type Task struct {
	Category    string            `yaml:"category"`
	Description string            `yaml:"description"`
	DependsOn   []string          `yaml:"depends_on"`
	Env         map[string]string `yaml:"env"`
	Quiet       bool              `yaml:"quiet"`
	Commands    []string          `yaml:"commands"`
}

type TemplatesConfig struct {
	Sprig bool `yaml:"sprig"`
}

type Config struct {
	App       AppConfig         `yaml:"app"`
	Env       map[string]string `yaml:"env"`
	Templates TemplatesConfig   `yaml:"templates"`
	Tasks     map[string]Task   `yaml:"tasks"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse yaml: %w", err)
	}

	if err := validateConfig(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

var (
	taskNamePattern = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
	envKeyPattern   = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	versionPattern  = regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[a-zA-Z0-9.]+)?$`)
	mainFilePattern = regexp.MustCompile(`^(\./)?[a-zA-Z0-9_/-]+\.go$`)
)

func validateConfig(cfg Config) error {
	var issues []string

	if strings.TrimSpace(cfg.App.Name) == "" {
		issues = append(issues, "app.name is required")
	}
	if strings.TrimSpace(cfg.App.Version) == "" {
		issues = append(issues, "app.version is required")
	} else if !versionPattern.MatchString(cfg.App.Version) {
		issues = append(issues, "app.version must be semver (e.g. v1.2.3)")
	}
	if strings.TrimSpace(cfg.App.Description) == "" {
		issues = append(issues, "app.description is required")
	}
	if strings.TrimSpace(cfg.App.MainFile) == "" {
		issues = append(issues, "app.main_file is required")
	} else if !mainFilePattern.MatchString(cfg.App.MainFile) {
		issues = append(issues, "app.main_file must be a .go path")
	}
	if len(cfg.App.Authors) == 0 {
		issues = append(issues, "app.authors must include at least one entry")
	}
	if strings.TrimSpace(cfg.App.RepoURL) == "" {
		issues = append(issues, "app.repo_url is required")
	} else if !isValidURL(cfg.App.RepoURL) {
		issues = append(issues, "app.repo_url must be a valid URL")
	}

	for key := range cfg.Env {
		if !envKeyPattern.MatchString(key) {
			issues = append(issues, fmt.Sprintf("env key %q is invalid", key))
		}
	}

	if len(cfg.Tasks) == 0 {
		issues = append(issues, "no tasks defined")
	}
	if cfg.Env == nil {
		cfg.Env = map[string]string{}
	}

	for name, task := range cfg.Tasks {
		if task.Env == nil {
			task.Env = map[string]string{}
			cfg.Tasks[name] = task
		}

		if !taskNamePattern.MatchString(name) {
			issues = append(issues, fmt.Sprintf("task name %q is invalid", name))
		}
		if strings.TrimSpace(task.Category) == "" {
			issues = append(issues, fmt.Sprintf("task %q: category is required", name))
		}
		if strings.TrimSpace(task.Description) == "" {
			issues = append(issues, fmt.Sprintf("task %q: description is required", name))
		}
		if len(task.Commands) == 0 {
			issues = append(issues, fmt.Sprintf("task %q: commands must not be empty", name))
		}
		if len(task.DependsOn) > 1 {
			issues = append(issues, fmt.Sprintf("task %q: depends_on must have at most one entry", name))
		}
		for key := range task.Env {
			if !envKeyPattern.MatchString(key) {
				issues = append(issues, fmt.Sprintf("task %q: env key %q is invalid", name, key))
			}
		}

		for _, dep := range task.DependsOn {
			if _, ok := cfg.Tasks[dep]; !ok {
				issues = append(issues, fmt.Sprintf("task %q: depends_on %q not found", name, dep))
			}
		}
	}

	if len(issues) > 0 {
		return fmt.Errorf("invalid tasks.yml:\n- %s", strings.Join(issues, "\n- "))
	}

	return nil
}

func isValidURL(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}
	return u.Scheme != "" && u.Host != ""
}
