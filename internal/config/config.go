package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	MainFile    string `yaml:"main_file"`
	Author      string `yaml:"author"`
	RepoURL     string `yaml:"repo_url"`
}

type Task struct {
	Category    string   `yaml:"category"`
	Description string   `yaml:"description"`
	DependsOn   []string `yaml:"depends_on"`
	Commands    []string `yaml:"commands"`
}

type Config struct {
	App     AppConfig       `yaml:"app"`
	Tasks   map[string]Task `yaml:"tasks"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse yaml: %w", err)
	}

	if len(cfg.Tasks) == 0 {
		return Config{}, fmt.Errorf("no tasks defined")
	}

	return cfg, nil
}
