package server

import (
	"context"

	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, ctx context.Context, log logger.ILogger, grpc *grpc.GrpcServer, db *services.DatabaseServer) error {
	lc.Append(fx.Hook {
		OnStart: func(_ context.Context) error {
			go func() {
				if err := grpc.Run(ctx); err != nil {
					log.Fatal("Error running grpc server: " + err.Error())
				}
			}()

			go func() {
				if err := db.Run(ctx); err != nil {
					log.Fatal("Error running database: " + err.Error())
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