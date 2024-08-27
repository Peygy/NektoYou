package gin

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type GinConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type GinServer struct {
	Engine 	*gin.Engine
	Config 	*GinConfig
	Log 	logger.ILogger
}

func NewGinServer(cfg *GinConfig, log logger.ILogger) *GinServer {
	log.Info("Gin engine is created")
	ginEngine := gin.Default()
	return &GinServer{ ginEngine, cfg, log }
}

func (s *GinServer) Run(ctx context.Context) error {
	address := s.Config.Host+s.Config.Port

	go func () {
		for {
			select {
			case <-ctx.Done():
				s.Log.Info("Shutting down gin on address: " + address)
				return
			}
		}
	} ()

	if err := s.Engine.Run(address); err != nil {
		s.Log.Fatal("Gin server can't be runned on address: " + address)
		return err
	}

	s.Log.Info("Gin server runned on address: " + address)
	return nil
}