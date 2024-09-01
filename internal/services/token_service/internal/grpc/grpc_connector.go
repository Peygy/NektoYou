package grpc

import (
	"context"
	"time"

	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_token"
	"github.com/peygy/nektoyou/internal/services/token_service/internal/jwt"
	"github.com/peygy/nektoyou/internal/services/token_service/internal/managers"
)

func InitTokenGrpcServer(
	server *grpc.GrpcServer,
	tokenManager jwt.ITokenManager,
	refreshManager managers.IRefreshManager,
	logger logger.ILogger) {
	grpcServer := &grpcServer{
		tokenManager:   tokenManager,
		refreshManager: refreshManager,
		log:            logger,
	}
	pb.RegisterCreateTokensPairServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedCreateTokensPairServiceServer

	tokenManager   jwt.ITokenManager
	refreshManager managers.IRefreshManager

	log logger.ILogger
}

func (s *grpcServer) CreateTokensPair(ctx context.Context, in *pb.CreateTokensPairRequest) (*pb.CreateTokensPairResponce, error) {
	rt, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	err = s.refreshManager.AddToken(in.Userid, rt)
	if err != nil {
		return nil, err
	}

	at, err := s.tokenManager.NewAccessToken(in.Userid, time.Minute*5, in.Roles...)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTokensPairResponce{AccessToken: at, RefreshToken: rt}, nil
}
