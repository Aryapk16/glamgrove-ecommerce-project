package main

import (
	"glamgrove/pkg/config"
	"glamgrove/pkg/di"
	"log"
)

func main() {

	config, Err := config.LoadConfig()
	if Err != nil {
		log.Fatal("cannot load config : ", Err)
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
