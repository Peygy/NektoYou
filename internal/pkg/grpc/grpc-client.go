package grpc

import (
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn *grpc.ClientConn
}

func NewGrpcClient(cfg *GrpcConfig, log logger.ILogger) (*GrpcClient, error) {
	conn, err := grpc.NewClient(cfg.Host + cfg.Port)
	if err != nil {
		log.Error("Error while create grpc client: " + err.Error())
		return nil, err
	}

	return &GrpcClient{ conn }, nil
}