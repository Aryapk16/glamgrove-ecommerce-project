package usecase

import (
	"context"
	"errors"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type OrderUseCase struct {
	OrderRepository interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository) service.OrderService {
	return &OrderUseCase{
		OrderRepository: repo,
	}
}
func (o *OrderUseCase) GetTotalAmount(c context.Context, userid uint) (float64, error) {
	var total_amount float64
	total_amount = 0
	cart, err := o.OrderRepository.GetTotalAmount(c, int(userid))
	if err != nil {
		return 0, err
	}

	for _, c := range cart {
		total_amount = total_amount + float64(c.Total)
	}
	return total_amount, nil
}
func (o *OrderUseCase) CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := o.OrderRepository.FindPaymentMethodById(c, uint(order.PaymentMethodID))

	if err != nil {
		return response.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderresp, err := o.OrderRepository.CreateOrder(c, order)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderresp, nil
}

func (o *OrderUseCase) UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := o.OrderRepository.FindPaymentMethodById(c, uporder.PaymentMethodID)

	if err != nil {
		return response.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderup, err := o.OrderRepository.UpdateOrderDetails(c, uporder)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderup, nil
}
func (o *OrderUseCase) ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error) {

	return o.OrderRepository.ListAllOrders(c, page, userId)

}

// List order for admin
func (o *OrderUseCase) GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error) {

	return o.OrderRepository.GetAllOrders(c, page)

}
func (o *OrderUseCase) DeleteOrder(c context.Context, order_id uint) error {
	err := o.OrderRepository.DeleteOrder(c, order_id)
	if err != nil {
		return err
	}
	return nil
}
func (o *OrderUseCase) PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error) {
	err1 := o.OrderRepository.FindOrder(c, order)
	if err1 != nil {
		return response.PaymentResponse{}, err1
	}
	method_Id, err := o.OrderRepository.FindPaymentMethodIdByOrderId(c, order.Order_Id)
	if err != nil {
		return response.PaymentResponse{}, err
	}
	if method_Id == 1 {
		order.Order_Status = "order confirmed"
	} else {
		order.Order_Status = "order confirmed payment pending"
	}
	paymentresp, err := o.OrderRepository.PlaceOrder(c, order)
	if err != nil {
		return response.PaymentResponse{}, err
	}
	return paymentresp, nil
}

func (o *OrderUseCase) FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error) {
	method_id, err := o.OrderRepository.FindPaymentMethodIdByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return method_id, nil
}

func (o *OrderUseCase) UpdateOrderStatus(c context.Context, order_id uint) (response.OrderResponse, error) {
	order_status := "Order confirmed"
	orderResp, err := o.OrderRepository.UpdateOrderStatus(c, order_id, order_status)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderResp, nil
}
func (o *OrderUseCase) FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error) {
	totalAmount, err := o.OrderRepository.FindTotalAmountByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return totalAmount, nil
}
func (o *OrderUseCase) ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error) {
	returnResp, err := o.OrderRepository.ReturnRequest(c, returnOrder)
	if err != nil {
		return response.ReturnResponse{}, err
	}
	return returnResp, nil
}

func (o *OrderUseCase) VerifyOrderID(c context.Context, id uint, orderid uint) error {
	err := o.OrderRepository.VerifyOrderID(c, id, orderid)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUseCase) GetAllPendingReturnRequest(c context.Context, page request.ReqPagination) ([]response.ReturnRequests, error) {

	returnRequests, err := o.OrderRepository.GetAllPendingReturnOrder(c, page)
	if err != nil {
		return returnRequests, err
	}

	return returnRequests, nil

}
