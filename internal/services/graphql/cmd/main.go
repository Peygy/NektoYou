package main

import (
	"go.uber.org/fx"
	"github.com/peygy/nektoyou/internal/services/graphql/server"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

func main () {
	fx.New(
		fx.Options(
			fx.Provide( // определяем экземпляры сервисов, e.x сервер
				logger.NewLogger,
			),
			fx.Invoke(server.RunServers), // определим маршруты
		),
	).Run()
}