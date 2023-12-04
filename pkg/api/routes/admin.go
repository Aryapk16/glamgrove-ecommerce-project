package routes

import (
	"glamgrove/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.Engine, adminHandler *handler.AdminHandler) {
	//login
	login := api.Group("/adminlogin")
	{
		login.POST("/", adminHandler.AdminLogin)

	}
}
