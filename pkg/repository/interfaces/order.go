package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type OrderRepository interface {
	CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error)
	UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error)
	DeleteOrder(c context.Context, order_id uint) error
	ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error)
	GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error)

	FindPaymentMethodById(c context.Context, method_id uint) (uint, error)
	FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error)
	GetTotalAmount(c context.Context, userid int) ([]domain.Cart, error)

	FindOrder(c context.Context, order domain.Order) error
	PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error)
	UpdateOrderStatus(c context.Context, order_id uint, order_status string) (response.OrderResponse, error)

	FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error)

	ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error)
	VerifyOrderID(c context.Context, id uint, orderid uint) error

	GetAllPendingReturnOrder(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error)
}
