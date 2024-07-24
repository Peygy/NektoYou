package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type GinConfig struct {
	Port string `yaml:"port"`
}

func NewGinServer(cfg *GinConfig, log logger.ILogger) *gin.Engine {
	log.Info("Gin engine is created")
	return gin.Default()
}

func RunGinServer(eng *gin.Engine, cfg *GinConfig, log logger.ILogger) error {
	if err := eng.Run(cfg.Port); err != nil {
		log.Fatal("Gin server can't be runned on port " + cfg.Port)
		return err
	}

	log.Info("Gin server runned on port " + cfg.Port)
	return nil
}