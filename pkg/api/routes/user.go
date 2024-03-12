package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, paymentHandler *handler.PaymentHandler, couponHandler *handler.CouponHandler, orderHandler *handler.OrderHandler) {

	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignup)
		signup.POST("/otp/verify", userHandler.VerifyOtp)

	}

	login := api.Group("/login")
	{
		login.POST("/", userHandler.LoginSubmit)
	}

	//forgotpassword
	forgotpass := api.Group("/sendotp")
	{
		forgotpass.POST("/", userHandler.SendOtpForgotPass)
	}

	verifyOrder := api.Group("/order")
	{
		verifyOrder.POST("/order/verify", productHandler.RazorpayVerify)
	}

	//api.POST("/order/verify", productHandler.RazorpayVerify)

	// Middleware
	api.Use(middleware.AuthenticateUser)
	{

		products := api.Group("/products")
		{
			products.GET("/brands", productHandler.GetAllCategory)
			products.GET("/", productHandler.ListProducts)
			//products.GET("/productitem", productHandler.GetProductItem)
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

		coupon := api.Group("/coupons")
		{
			coupon.GET("/list", couponHandler.ListAllCoupons)
		}

		order := api.Group("/order")
		{
			order.POST("/createOrder", orderHandler.CreateOrder)
			order.PUT("/updateOrder", orderHandler.UpdateOrder)
			order.GET("/listOrder", orderHandler.ListAllOrders)
			order.DELETE("/cancelOrder", orderHandler.CancelOrder)
			order.POST("/placeOrder", orderHandler.PlaceOrder)
			order.POST("/payment", orderHandler.CheckOut)

			//order.GET("/print", orderHandler.PrintInvoice)
		}
		Return := api.Group("/return")
		{
			Return.POST("/product", orderHandler.ReturnOrder)
		}

	}
}
