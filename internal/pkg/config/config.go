package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfig[T any](filePath string) (*T, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cfg := new(T)

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
