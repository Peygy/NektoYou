package internal

import (
	"context"

	"github.com/peygy/nektoyou/internal/pkg/grpc"
	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
)

func InitAuthGrpcServer(server *grpc.GrpcServer) {
	pb.RegisterSignInServiceServer(server.Engine, &grpcServer{})
}

type grpcServer struct {
	pb.UnimplementedSignInServiceServer
}

func (s *grpcServer) GeneratePairOfTokens(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponce, error) {
	return &pb.SignInResponce{}, nil
}