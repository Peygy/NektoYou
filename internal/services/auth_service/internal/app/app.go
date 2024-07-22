package app

import (
	"api-gateway/internal/routes"
)

func Launch(port string) error {
	router := routes.CreateRoutes()

	err := router.Run(port)
	return err
}
