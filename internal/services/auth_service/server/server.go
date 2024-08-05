package server

import (
	"context"

	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, ctx context.Context, log logger.ILogger, gin *gin.GinServer, grpc *grpc.GrpcServer) error {
	lc.Append(fx.Hook {
		OnStart: func(_ context.Context) error {
			/*go func() {
				if err := gin.Run(ctx); err != nil {
					log.Fatal("Error running gin server: " + err.Error())
				}
			}()*/
			go func() {
				if err := grpc.Run(ctx); err != nil {
					log.Fatal("Error running grpc server: " + err.Error())
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