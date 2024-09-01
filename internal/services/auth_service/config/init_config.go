package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = "./config/config.dev.yml"

type AuthConfig struct {
	GrpcServer     *grpc.GrpcServerConfig `yaml:"grpc-server"`
	DatabaseConfig *DatabaseConfig        `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

func NewAuthConfig() (*AuthConfig, *grpc.GrpcServerConfig, *DatabaseConfig, error) {
	cfg, err := config.NewConfig[AuthConfig](configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.DatabaseConfig, nil
}
