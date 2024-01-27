package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {

	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignup)
		signup.POST("/otp/verify", userHandler.VerifyOtp)

	}

	login := api.Group("/login")
	{
		login.POST("/", userHandler.LoginSubmit)
		//login.POST("/otp-verify", userHandler.UserOTPVerify)
	}
	
	//forgotpassword
	forgotpass := api.Group("/sendotp")
	{
		forgotpass.POST("/", userHandler.SendOtpForgotPass)
	}

	// Middleware
	api.Use(middleware.AuthenticateUser)
	{

		products := api.Group("/products")
		{
			products.GET("/brands", productHandler.GetAllCategory)
			products.GET("/", productHandler.ListProducts)
		}
		profile := api.Group("/profile")
		{
			profile.POST("/add-address", userHandler.AddAddress)
			profile.GET("/get-address", userHandler.GetAllAddress)
			profile.DELETE("/delete-address/:adressId", userHandler.DeleteAddress)
			profile.PUT("/edit-address", userHandler.UpdateAddress)
			profile.GET("/", userHandler.Profile)

		}
		cart := api.Group("/cart")
		{
			cart.POST("/add", userHandler.AddToCart)
			cart.GET("/get", userHandler.GetcartItems)
			cart.PUT("/update", userHandler.UpdateCart)
			cart.DELETE("/delete", userHandler.DeleteCartItem)
		}

		order := api.Group("/order")
		{
			order.POST("/createOrder", orderHandler.CreateOrder)
			order.PUT("/updateOrder", orderHandler.UpdateOrder)
			order.GET("/listOrder", orderHandler.ListAllOrders)
			order.DELETE("/cancelOrder", orderHandler.CancelOrder)
			order.POST("/placeOrder", orderHandler.PlaceOrder)
			//order.POST("/payment", orderHandler.CheckOut)
		}
		Return := api.Group("/return")
		{
			Return.POST("/product", orderHandler.ReturnOrder)
		}

	}
}
