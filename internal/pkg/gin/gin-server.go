package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type GinConfig struct {
	Port 	string
}

func NewGinServer(cfg *GinConfig, log logger.ILogger) *gin.Engine {
	return gin.Default()
}

func RunGinServer(eng *gin.Engine, cfg *GinConfig, log logger.ILogger) error {
	err := eng.Run(cfg.Port)
	return err
}