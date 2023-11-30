package main

import (
	"glamgrove/pkg/config"
	"glamgrove/pkg/di"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Error to load the config: ", err)
	}
	server, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatal("Failed to initialize the api: ", err)
	}
	if server.Start(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
