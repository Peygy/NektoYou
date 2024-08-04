package config

import (
	"os"

	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Gin 		*gin.GinConfig         `yaml:"gin"`
	GrpcServer	*grpc.GrpcServerConfig `yaml:"grpc-server,omitempty"`
	GrpcClient	*grpc.GrpcClientConfig `yaml:"grpc-client,omitempty"`
}

func InitConfig(filePath string) (*Config, *gin.GinConfig, *grpc.GrpcServerConfig, *grpc.GrpcClientConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return cfg, cfg.Gin, cfg.GrpcServer, cfg.GrpcClient, nil
}
