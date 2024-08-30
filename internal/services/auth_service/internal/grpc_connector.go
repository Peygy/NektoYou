package internal

import (
	"context"
	"time"

	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/managers"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services/jwt"
)

func InitAuthGrpcServer(
	server *grpc.GrpcServer,
	tokenManager jwt.ITokenManager,
	roleManager managers.IRoleManager,
	userManager managers.IUserManager,
	refreshManager managers.IRefreshManager,
	logger logger.ILogger) {
	grpcServer := &grpcServer{
		tokenManager:   tokenManager,
		roleManager:    roleManager,
		userManager:    userManager,
		refreshManager: refreshManager,
		log:            logger,
	}
	pb.RegisterSignUpServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedSignUpServiceServer

	tokenManager   jwt.ITokenManager
	roleManager    managers.IRoleManager
	userManager    managers.IUserManager
	refreshManager managers.IRefreshManager

	log logger.ILogger
}

func (s *grpcServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponce, error) {
	user := managers.UserRecord{UserName: in.Username, Password: in.Password}
	userId, err := s.userManager.InsertUser(user)
	if err != nil {
		return nil, err
	}

	err = s.roleManager.AddRolesToUser(userId, "user")
	if err != nil {
		return nil, err
	}

	rt, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	err = s.refreshManager.AddToken(userId, rt)
	if err != nil {
		return nil, err
	}

	at, err := s.tokenManager.NewAccessToken(userId, time.Minute*5)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponce{AccessToken: at, RefreshToken: rt}, nil
}
