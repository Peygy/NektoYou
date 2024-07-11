package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CreateRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/:id", hi)

	return router
}

func hi(c *gin.Context) {
	log.Print(c.Param("id"))
}
