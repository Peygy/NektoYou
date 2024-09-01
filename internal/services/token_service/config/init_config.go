package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = "./config/config.dev.yml"

type TokenServiceConfig struct {
	GrpcServer     *grpc.GrpcServerConfig `yaml:"grpc-server"`
	TokenConfig    *TokenConfig           `yaml:"token-config"`
	DatabaseConfig *DatabaseConfig        `yaml:"database"`
}

type TokenConfig struct {
	SecretKey string `yaml:"secretKey"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

func NewTokenServiceConfig() (*TokenServiceConfig, *grpc.GrpcServerConfig, *TokenConfig, *DatabaseConfig, error) {
	cfg, err := config.NewConfig[TokenServiceConfig](configPath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return cfg, cfg.GrpcServer, cfg.TokenConfig, cfg.DatabaseConfig, nil
}
