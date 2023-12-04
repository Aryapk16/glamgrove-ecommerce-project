package http

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler) *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())

	//two main routes `\` -> user ; `\admin`-> admin
	routes.UserRoutes(engine, userHandler)
	routes.AdminRoutes(engine, adminHandler)
	

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() {

	s.engine.Run(":8000")
}
