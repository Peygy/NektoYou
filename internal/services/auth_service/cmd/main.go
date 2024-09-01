package main

import (
	"github.com/peygy/nektoyou/internal/pkg/context"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/data"
	grpcConn "github.com/peygy/nektoyou/internal/services/auth_service/internal/grpc"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/managers"
	"github.com/peygy/nektoyou/internal/services/auth_service/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewAuthConfig,
				logger.NewLogger,
				context.NewContext,
				grpc.NewGrpcServer,

				data.NewDatabaseConnection,
				managers.NewRoleManager,
				managers.NewUserManager,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitAuthGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
