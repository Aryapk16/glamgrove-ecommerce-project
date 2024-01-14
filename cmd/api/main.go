package main

import (
	"glamgrove/pkg/config"
	"glamgrove/pkg/di"
	"log"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error loading the env file")
	// }
	config, Err := config.LoadConfig()
	if Err != nil {
		log.Fatal("cannot load config : ", Err)
	}
	server, diErr := di.InitializeApi(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
