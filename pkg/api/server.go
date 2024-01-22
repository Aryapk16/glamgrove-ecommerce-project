package http

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	//Calling routes
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, orderHandler, paymentHandler)
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, orderHandler, paymentHandler)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
