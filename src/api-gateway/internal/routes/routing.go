package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/login", login)
	router.POST("/register", registration)

	privateRouter := router.Group("/api")
	privateRouter.Use(jwtTokenCheck)

	return router
}
