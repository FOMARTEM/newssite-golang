package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadConfig(pathToFile string) (*Config, error) {
	filename, err := filepath.Abs(pathToFile)
	if err != nil {
		return nil, err
	}

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
