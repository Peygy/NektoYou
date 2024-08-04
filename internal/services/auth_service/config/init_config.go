package config

import (
	"github.com/peygy/nektoyou/internal/pkg/config"
	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

const configPath = ".\\config\\config.dev.yml"

func NewConfig() (*config.Config, *gin.GinConfig, *grpc.GrpcServerConfig, *grpc.GrpcClientConfig, error) {
	cfg := config.InitConfig(configPath)

	return cfg, cfg.Gin, cfg.GrpcServer, cfg.GrpcClient, nil
}
