package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type OrderService interface {
	GetTotalAmount(c context.Context, userid uint) (float64, error)
	CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error)
	ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error)
	UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error)

	GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error)
	DeleteOrder(c context.Context, order_id uint) error
	PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error)
	FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error)

	FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error)
	FindPhoneEmailByUserId(c context.Context, usr_id int) (response.PhoneEmailResp, error)

	GetRazorpayOrder(c context.Context, userID uint, razorPay request.RazorPayReq) (response.ResRazorpayOrder, error)
	UpdateStatusRazorpay(c context.Context, order_id uint) (response.OrderResponse, error)
	UpdateOrderStatus(c context.Context, order_id uint) (response.OrderResponse, error)

	ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error)
	VerifyOrderID(c context.Context, id uint, orderid uint) error

	SalesReport(c context.Context, page request.ReqPagination, salesData request.ReqSalesReport) ([]response.SalesReport, error)

	GetAllPendingReturnRequest(c context.Context, page request.ReqPagination) ([]response.ReturnRequests, error)
}
