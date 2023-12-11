package routes

import (
	"glamgrove/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, user *handler.UserHandler) {

	router.POST("/login", user.Login)
	router.POST("/signup", user.SignUp)
	router.POST("/verifyOTP", user.VerifyOTP)

	//api := router.Group("/", middleware.Authentication)
	//
	//api.GET("/", user.Home)
}
