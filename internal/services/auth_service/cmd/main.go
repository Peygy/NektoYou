package main

import (
	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
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
			),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
