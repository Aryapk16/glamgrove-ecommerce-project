package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, user *handler.UserHandler, productHandler *handler.ProductHandler) {

	router.POST("/login", user.Login)
	router.POST("/signup", user.SignUp)
	router.POST("/verifyOTP", user.VerifyOTP)

	router.Use(middleware.AuthenticateUser)
	{
		listproducts := router.Group("/product")
		{
			listproducts.GET("/getallproducts", productHandler.GetAllProducts)
			listproducts.GET("/search", productHandler.SearchProduct)

		}

	}

}
