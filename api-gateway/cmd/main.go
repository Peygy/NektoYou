package main

import (
	config "api-gateway/configs"
	"api-gateway/internal/app"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := app.Launch(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
