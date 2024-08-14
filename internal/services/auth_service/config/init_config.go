package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = ".\\config\\config.dev.yml"

func NewConfig() (*config.Config, *grpc.GrpcServerConfig, error) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		return nil, nil, err
	}

	return cfg, cfg.GrpcServer, nil
}
