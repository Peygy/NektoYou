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
		connection := val.Host + val.Port
		conn, err := grpc.NewClient(connection, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Errorf("Can't create grpc client %s on connection %s with error: %v", val.Name, connection, err)
			return nil, err
		}

		connPull.Services = append(connPull.Services, GrpcService{val.Name, conn})
		log.Infof("To grpc pull is added new service: %s on connection %s", val.Name, connection)
	}

	return connPull, nil
}
