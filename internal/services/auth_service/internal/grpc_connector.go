package internal

import (
	"context"

	"github.com/peygy/nektoyou/internal/pkg/grpc"
	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
)

func InitAuthGrpcServer(server *grpc.GrpcServer) {
	pb.RegisterTokensGeneraterServer(server.Engine, &grpcServer{})
}

type grpcServer struct {
	pb.UnimplementedTokensGeneraterServer
}

func (s *grpcServer) GenerateAuthTokens(ctx context.Context, in *pb.AuthTokensRequest) (*pb.AuthTokensResponce, error) {
	return &pb.AuthTokensResponce{}, nil
}