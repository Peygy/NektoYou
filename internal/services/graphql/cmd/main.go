package main

import (
	"go.uber.org/fx"
	"github.com/peygy/nektoyou/internal/services/graphql/server"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

func main () {
	fx.New(
		fx.Options(
			fx.Provide(
				logger.NewZapLogger(),
			),
			fx.Invoke(server.RunServers),
		),
	).Run()
}