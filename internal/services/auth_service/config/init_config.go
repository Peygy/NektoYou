package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = ".\\config\\config.dev.yml"

func NewConfig() (*config.Config, *gin.GinConfig, *grpc.GrpcServerConfig, error) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, cfg.Gin, cfg.GrpcServer, nil
}
