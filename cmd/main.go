package main

import (
	"log"
	"web/config"
	"web/pkg/server"
	"github.com/spf13/viper"
)

func main() {	
	if err := config.Init(); err != nil{
		log.Fatalf("%s", err.Error())
	}

	if err := server.Launch(viper.GetString("port")); err != nil{
		log.Fatalf("%s", err.Error())
	}
}