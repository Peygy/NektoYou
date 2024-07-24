package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/peygy/nektoyou/internal/pkg/config"
	gin_server "github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, log logger.ILogger, eng *gin.Engine, cfg *config.Config) error {
	lc.Append(fx.Hook {
		OnStart: func(_ context.Context) error {
			go func() {
				if err := gin_server.RunGinServer(eng, cfg.Gin, log); err != nil {
					log.Fatal("Error running gin server")
				}
			}()

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Info("All servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}