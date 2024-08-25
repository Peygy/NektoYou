package services

import (
	"context"

	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
)

type grpcServer struct {
	pb.UnimplementedTokensGeneraterServer
}

func (s *grpcServer) GenerateAuthTokens(ctx context.Context, in *pb.AuthTokensRequest) (*pb.AuthTokensResponce, error) {

	return &pb.AuthTokensResponce{}, nil
}

func generateJwtAccessToken() {

}

func generateRefreshToken() {
	
}