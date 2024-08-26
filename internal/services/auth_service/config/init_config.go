package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = ".\\config\\config.dev.yml"

type AuthConfig struct {
	GrpcServer   *grpc.GrpcServerConfig `yaml:"grpc-server"`
	TokenManager *TokenManagerConfig    `yaml:"token-manager"`
}

type TokenManagerConfig struct {
	SecretKey string `yaml:"secretKey"`
}

func NewAuthConfig() (*AuthConfig, *grpc.GrpcServerConfig, *TokenManagerConfig, error) {
	cfg, err := config.NewConfig[AuthConfig](configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.TokenManager, nil
}
