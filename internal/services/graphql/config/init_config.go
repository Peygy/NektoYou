package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/gin"
)

const configPath = ".\\config\\config.dev.yml"

func NewConfig() (*config.Config, *gin.GinConfig, error) {
	return config.InitConfig(configPath)
}
