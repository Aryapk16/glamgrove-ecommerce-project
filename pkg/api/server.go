package http

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler,
) *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())

	//two main routes `\` -> user ; `\admin`-> admin
	routes.UserRoutes(engine, userHandler,productHandler)
	routes.AdminRoutes(engine, adminHandler,productHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() {

	s.engine.Run(":8000")
}
