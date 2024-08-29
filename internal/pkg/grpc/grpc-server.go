package grpc

import (
	"context"
	"net"
	"time"

	"github.com/peygy/nektoyou/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type GrpcServer struct {
	Engine *grpc.Server
	config *GrpcServerConfig
	log    logger.ILogger
}

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func NewGrpcServer(cfg *GrpcServerConfig, log logger.ILogger) *GrpcServer {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
	)

	return &GrpcServer{grpcServer, cfg, log}
}

func (s *GrpcServer) Run(ctx context.Context) error {
	address := s.config.Host + s.config.Port

	listen, err := net.Listen("tcp", address)
	if err != nil {
		s.log.Fatalf("Grpc server can't be runned on address: %s", address)
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.log.Infof("Shutting down grpc server on address: ", address)
				s.shutdown()
				s.log.Info("Grpc exited properly")
				return
			}
		}
	}()

	s.log.Infof("Grpc server is listening on address: %s", address)

	err = s.Engine.Serve(listen)
	if err != nil {
		s.log.Fatalf("Grpc server can't server connection with error: %v", err)
		return err
	}

	return nil
}

func (s *GrpcServer) shutdown() {
	s.Engine.Stop()
	s.Engine.GracefulStop()
}
