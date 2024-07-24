package config

import (
	"os"

	"github.com/peygy/nektoyou/internal/pkg/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Gin *gin.GinConfig `yaml:"gin"`
}

func InitConfig(filePath string) (*Config, *gin.GinConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, nil, err
	}

	return cfg, cfg.Gin, nil
}
