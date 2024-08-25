package services

import (
	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
)

func InitAuthGrpcServer(server *grpc.GrpcServer) {
	pb.RegisterTokensGeneraterServer(server.Engine, &grpcServer{})
}
