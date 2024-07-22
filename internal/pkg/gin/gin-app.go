package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type GinConfig struct {
	Port string
}

func NewGinServer(cfg *GinConfig, log logger.ILogger) {
	r := gin.Default()
	r.Run(cfg.Port)
}
