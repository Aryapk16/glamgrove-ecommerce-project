package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler) {
	//login
	login := api.Group("/login")
	{
		login.POST("/", adminHandler.AdminLogin)

	}
	//middleware

	api.Use(middleware.AuthenticateAdmin)
	{
		user := api.Group("/users")
		{
			user.GET("/", adminHandler.ListUsers)
			user.PATCH("/block", adminHandler.BlockUnBlockUser)
			user.PATCH("/unblock", adminHandler.BlockUnBlockUser)
		}

		brand := api.Group("/brands")
		{
			brand.GET("/get", productHandler.GetAllCategory)
			brand.POST("/add", productHandler.AddCategory)
		}

		product := api.Group("/products")
		{
			product.GET("/list", productHandler.ListProducts)
			product.POST("/add", productHandler.AddProduct)
			product.PUT("/update", productHandler.UpdateProduct)
			product.DELETE("/delete", productHandler.DeleteProduct)
			product.POST("/product-item", productHandler.AddProductItem)
			product.GET("/product-item/:product_id", productHandler.GetProductItem)
		}
	}
}
