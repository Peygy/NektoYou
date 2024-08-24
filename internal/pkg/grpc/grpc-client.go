package grpc

import (
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientConfig struct {
	Services []struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"services"`
}

type GrpcService struct {
	Name string
	Conn *grpc.ClientConn
}

type GrpcPull struct {
	Services []GrpcService
}

func NewGrpcClient(cfg *GrpcClientConfig, log logger.ILogger) (*GrpcPull, error) {
	connPull := new(GrpcPull)
	for _, val := range cfg.Services {
		conn, err := grpc.NewClient(val.Host+val.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Error("Error while create grpc server" + val.Name + " connection: " + err.Error())
			return nil, err
		}

		connPull.Services = append(connPull.Services, GrpcService{val.Name, conn})
		log.Info("To grpc pull is added new service: " + val.Name)
	}

	return connPull, nil
}
