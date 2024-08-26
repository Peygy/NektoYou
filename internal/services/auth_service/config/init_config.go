package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services"
)

const configPath = ".\\config\\config.dev.yml"

type AuthConfig struct {
	GrpcServer   *grpc.GrpcServerConfig `yaml:"grpc-server"`
	TokenManager *services.Manager      `yaml:"token-manager"`
}

func NewAuthConfig() (*AuthConfig, *grpc.GrpcServerConfig, *services.Manager, error) {
	cfg, err := config.NewConfig[AuthConfig](configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.TokenManager, nil
}
