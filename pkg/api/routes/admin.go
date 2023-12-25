package routes

import (
	"glamgrove/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.Engine, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler) {
	//login
	login := api.Group("/adminlogin")
	{
		login.POST("/", adminHandler.AdminLogin)

	}
	category := api.Group("/category")
	{
		category.POST("/", productHandler.SaveCategory)
		category.GET("/allcategory", productHandler.GetAllCategory)
	}
	products := api.Group("/products")
	{
		products.POST("/", productHandler.SaveProduct)
		products.GET("/getallproducts", productHandler.GetAllProducts)

	}
	user := api.Group("/users")
	{
		user.GET("/", adminHandler.GetAllUsers)
		user.PATCH("/blockuser", adminHandler.BlockUser)
		user.PATCH("/unblockuser", adminHandler.UnBlockUser)
	}
}
