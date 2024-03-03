package http

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	paymentHandler *handler.PaymentHandler,
	couponHandler *handler.CouponHandler,
	orderHandler *handler.OrderHandler,
	imageHandler *handler.ImageHandler) *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("views/*.html")

	// Get swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Calling routes
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, couponHandler, orderHandler, paymentHandler, imageHandler)
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, paymentHandler, couponHandler, orderHandler)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
