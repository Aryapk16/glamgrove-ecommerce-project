package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, admin *handler.AdminHandler) {

	router.POST("admin/login", admin.Login)

	api := router.Group("/admin", middleware.Authentication)

	api.GET("/alluser", admin.Allusers)
// 	api.GET("/add-product", admin.AddCategoryGET)
// 	api.POST("/add-product", admin.AddCategoryPOST)
 }
