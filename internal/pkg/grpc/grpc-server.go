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
	Config *GrpcServerConfig
	Log    logger.ILogger
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
	listen, err := net.Listen("tcp", s.Config.Port)
	if err != nil {
		s.Log.Fatal("Grpc server can't be runned on port " + s.Config.Port)
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.Log.Info("Shutting down Grpc on port: " + s.Config.Port)
				s.shutdown()
				s.Log.Info("Grpc exited properly")
				return
			}
		}
	}()

	s.Log.Info("Grpc server is listening on port: " + s.Config.Port)

	err = s.Engine.Serve(listen)
	if err != nil {
		s.Log.Error("Grpc server error: " + err.Error())
	}

	return nil
}

func (s *GrpcServer) shutdown() {
	s.Engine.Stop()
	s.Engine.GracefulStop()
}
