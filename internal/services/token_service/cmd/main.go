package main

import (
	"github.com/peygy/nektoyou/internal/pkg/context"
	"github.com/peygy/nektoyou/internal/pkg/database/postgres"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/token_service/config"
	"github.com/peygy/nektoyou/internal/services/token_service/internal/data"
	grpcConn "github.com/peygy/nektoyou/internal/services/token_service/internal/grpc"
	"github.com/peygy/nektoyou/internal/services/token_service/internal/jwt"
	"github.com/peygy/nektoyou/internal/services/token_service/internal/managers"
	"github.com/peygy/nektoyou/internal/services/token_service/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewTokenServiceConfig,
				logger.NewLogger,
				context.NewContext,
				grpc.NewGrpcServer,
				postgres.NewDatabaseConnection,

				managers.NewRefreshManager,
				jwt.NewTokenManager,
			),
			fx.Invoke(data.InitDatabaseSchema),
			fx.Invoke(grpcConn.InitTokenGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
