package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {

	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignup)

	}

	login := api.Group("/login")
	{
		login.POST("/", userHandler.LoginSubmit)
		login.POST("/otp-verify", userHandler.UserOTPVerify)
	}

	// Middleware
	api.Use(middleware.AuthenticateUser)
	{

		products := api.Group("/products")
		{
			products.GET("/brands", productHandler.GetAllCategory)
			products.GET("/", productHandler.ListProducts)
		}
	}
}
