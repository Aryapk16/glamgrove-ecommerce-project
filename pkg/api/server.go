package http

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) *ServerHTTP {
	engine := gin.New()

	//Calling routes
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler)
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
