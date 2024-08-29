package context

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

func NewContext(log logger.ILogger) context.Context {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Context is canceled!")
				cancel()
				return
			}
		}
	}()

	return ctx
}
