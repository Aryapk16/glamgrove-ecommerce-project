package http

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"
	"glamgrove/pkg/api/routes"

	"github.com/gin-contrib/cors"
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
	// corsConfig := cors.DefaultConfig()
	// //corsConfig.AllowOrigins = []string{"http://localhost:8000"} // Replace "yourport" with your actual port
	// corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	// corsConfig.AllowHeaders = []string{"Origin", "Content-Type"}
	// corsConfig.AllowAllOrigins = true

	// Apply CORS middleware with custom configuration
	//engine.Use(cors.New(corsConfig))

	//engine.Use(cors.Default())

	engine.LoadHTMLGlob("views/*.html")

	// Get swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Calling routes
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, couponHandler, orderHandler, paymentHandler, imageHandler)
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, paymentHandler, couponHandler, orderHandler)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow requests from all origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	sh.engine.Use(middleware.HandleOptionsRequest)
	sh.engine.Use(cors.New(config))
	sh.engine.Run(":8000")
}
