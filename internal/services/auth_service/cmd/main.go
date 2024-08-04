package main

import (
	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/grpc_server/services"
	"github.com/peygy/nektoyou/internal/services/auth_service/server"
	"go.uber.org/fx"
)

func main () {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewConfig,
				logger.NewLogger,
				gin.NewGinServer,
				grpc.NewGrpcServer,
			),
			fx.Invoke(services.InitAuthGrpcServer),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
