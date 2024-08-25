package services

import (
	"context"

	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

type grpcServer struct {
	pb.UnimplementedGreeterServer
}

func InitAuthGrpcServer(server *grpc.GrpcServer) {
	pb.RegisterGreeterServer(server.Engine, &grpcServer{})
}

func (s *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponce, error) {
	return &pb.HelloResponce{Message: "Hello again " + in.GetWord()}, nil
}