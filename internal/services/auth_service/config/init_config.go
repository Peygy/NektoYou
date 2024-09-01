package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/database/postgres"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = "./config/config.dev.yml"

type AuthConfig struct {
	GrpcServer     *grpc.GrpcServerConfig   `yaml:"grpc-server"`
	DatabaseConfig *postgres.PostgresConfig `yaml:"database"`
}

func NewAuthConfig() (*AuthConfig, *grpc.GrpcServerConfig, *postgres.PostgresConfig, error) {
	cfg, err := config.NewConfig[AuthConfig](configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.DatabaseConfig, nil
}
