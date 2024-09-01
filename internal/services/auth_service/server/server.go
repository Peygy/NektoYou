package server

import (
	"context"

	"github.com/peygy/nektoyou/internal/pkg/database/postgres"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, ctx context.Context, log logger.ILogger, grpc *grpc.GrpcServer, db postgres.IDatabaseServer) error {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := grpc.Run(ctx); err != nil {
					log.Fatalf("Error running grpc server: %v", err)
				}
			}()

			go func() {
				if err := db.Run(ctx); err != nil {
					log.Fatalf("Error running database: %v", err)
				}
			}()

			log.Info("Services are launched")
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Info("Logger buffer is flushed...")
			log.Sync()
			log.Info("All servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}
