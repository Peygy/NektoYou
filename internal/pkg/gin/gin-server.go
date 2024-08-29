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
	Engine *gin.Engine
	config *GinConfig
	log    logger.ILogger
}

func NewGinServer(cfg *GinConfig, log logger.ILogger) *GinServer {
	log.Info("Gin engine is created")
	ginEngine := gin.Default()
	return &GinServer{ginEngine, cfg, log}
}

func (s *GinServer) Run(ctx context.Context) error {
	address := s.config.Host + s.config.Port

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.log.Infof("Gin shutting down gin on address: %s", address)
				return
			}
		}
	}()

	if err := s.Engine.Run(address); err != nil {
		s.log.Fatalf("Gin server can't be runned on address %s with error %v", address, err)
		return err
	}

	s.log.Infof("Gin server runned on address: %s", address)
	return nil
}
