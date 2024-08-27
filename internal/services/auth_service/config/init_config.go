package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = "./config/config.dev.yml"

type AuthConfig struct {
	GrpcServer     *grpc.GrpcServerConfig `yaml:"grpc-server"`
	TokenManager   *TokenManagerConfig    `yaml:"token-manager"`
	DatabaseConfig *DatabaseConfig        `yaml:"database"`
}

type TokenManagerConfig struct {
	SecretKey string `yaml:"secretKey"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

func NewAuthConfig() (*AuthConfig, *grpc.GrpcServerConfig, *TokenManagerConfig, *DatabaseConfig, error) {
	cfg, err := config.NewConfig[AuthConfig](configPath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.TokenManager, cfg.DatabaseConfig, nil
}
