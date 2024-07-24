package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
)

const configPath = ".\\config\\config.dev.yml"

func NewConfig() {
	config.InitConfig(configPath)
}
