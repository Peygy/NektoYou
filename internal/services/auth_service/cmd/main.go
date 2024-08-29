package main

import (
	"github.com/peygy/nektoyou/internal/pkg/context"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/managers"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services/data"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services/jwt"
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

				jwt.NewTokenManager,
				data.NewDatabaseConnection,

				managers.NewRoleManager,
				managers.NewRefreshManager,
				managers.NewUserManager,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(internal.InitAuthGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
