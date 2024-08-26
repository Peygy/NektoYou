package config

import (
	"os"

	/*"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"*/
	"gopkg.in/yaml.v3"
)

/*type GrpcServerConfig struct {
	GrpcServer *grpc.GrpcServerConfig `yaml:"grpc-server,omitempty"`
}

type GrpcClientConfig struct {
	Gin        *gin.GinConfig         `yaml:"gin"`
	GrpcClient *grpc.GrpcClientConfig `yaml:"grpc-client,omitempty"`
}*/

/* Поменять на функцию NewConfig
func InitConfig(filePath string) (*Config, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}*/

func NewConfig[T any](filePath string, configStruct T) (*T, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cfg := new(T)

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
