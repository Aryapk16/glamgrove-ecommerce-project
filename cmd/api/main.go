package main

import (
	"glamgrove/docs"
	"glamgrove/pkg/config"
	"glamgrove/pkg/di"
	"log"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Go + Gin E-Commerce API Glamgrove
// @version 1.0.0
// @description Glamgrove is an E-commerce platform for purchasing high quality beauty products
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
// @query.collection.format multi

func main() {
	docs.SwaggerInfo.Title = "glamgrove"
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
