package routes

import (
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, couponHandler *handler.CouponHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {
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
			user.GET("/return-orders", adminHandler.GetAllReturnOrder)
			user.PATCH("/return-orders/approval", adminHandler.ApproveReturnOrder)
		}

		brand := api.Group("/brands")
		{
			brand.GET("/get", productHandler.GetAllCategory)
			brand.POST("/add", productHandler.AddCategory)
		}
		order := api.Group("/order")
		{
			order.GET("/listOrder", orderHandler.GetAllOrders)
		}

		product := api.Group("/products")
		{
			product.GET("/list", productHandler.ListProducts)
			product.POST("/add", productHandler.AddProduct)
			product.PUT("/update", productHandler.UpdateProduct)
			product.DELETE("/delete", productHandler.DeleteProduct)
			product.POST("/addimage", productHandler.AddImage)
			product.POST("/product-item", productHandler.AddProductItem)
			product.GET("/product-item/:product_id", productHandler.GetProductItem)
			product.POST("/additemimage", productHandler.AddItemImage)

		}

		coupons := api.Group("/coupons")
		{
			coupons.GET("/list", couponHandler.ListAllCoupons)
			coupons.POST("/create", couponHandler.CreateNewCoupon)
			coupons.DELETE("/invalid", couponHandler.MakeCouponInvalid)
			coupons.PUT("reactivate", couponHandler.ReActivateCoupon)
		}
		paymentmethod := api.Group("/paymentmethod")
		{
			paymentmethod.POST("/add", paymentHandler.AddpaymentMethod)
			paymentmethod.GET("/view", paymentHandler.GetPaymentMethods)
			paymentmethod.PUT("/update", paymentHandler.UpdatePaymentMethod)
			paymentmethod.DELETE("/delete", paymentHandler.DeleteMethod)
		}

		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/", adminHandler.DashBoard)
			dashboard.GET("/filteredSalesReport", adminHandler.FilteredSalesReport)
			dashboard.GET("/salesReport", orderHandler.SalesReport)
			dashboard.GET("/salesdata", productHandler.Statistics)
			dashboard.GET("/getallorders", orderHandler.ListAllOrders)

		}

	}
}
