package repository

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB: db}
}

// Get total amount
func (o *OrderDatabase) GetTotalAmount(c context.Context, userid int) ([]domain.Cart, error) {
	var cart []domain.Cart
	query := `select * from carts where user_id=?`
	err := o.DB.Raw(query, userid).Scan(&cart).Error
	if err != nil {
		return []domain.Cart{}, errors.New("failed to find cart items")
	}
	return cart, nil
}

// Insert into order table
func (o *OrderDatabase) CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error) {
	var orderdetails response.OrderResponse
	err := o.DB.Create(&order).Error
	if err != nil {
		return response.OrderResponse{}, errors.New("failed to place order")
	}
	query := `select o.order_id, o.total_amount, o.payment_status, o.order_status, o.delivery_status, o.address_id, p.payment_method from orders as o left join payment_methods as p on o.payment_method_id = p.id where o.order_id = ?`
	err1 := o.DB.Raw(query, order.Order_Id).Scan(&orderdetails).Error
	if err1 != nil {
		return response.OrderResponse{}, errors.New("failed to display order details")
	}
	return orderdetails, nil
}

// Find payment method by ID
func (o *OrderDatabase) FindPaymentMethodById(c context.Context, method_id uint) (uint, error) {
	var payment_methods domain.PaymentMethod
	err := o.DB.Raw("SELECT * FROM payment_methods WHERE id=?").Error
	if err != nil {
		return 0, errors.New("Failed to find payment method")
	}
	return payment_methods.ID, nil
}

// Update order details
func (o *OrderDatabase) UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error) {
	var order domain.Order
	var orderdetails response.OrderResponse
	query := `update orders set payment_method_id=?, address_id=?, payment_status=?, delivery_status=?  where order_id=?`
	err := o.DB.Raw(query, uporder.PaymentMethodID, uporder.Address_Id, uporder.Payment_Status, uporder.DeliveryStatus, uporder.Order_Id).Scan(&order).Error
	if err != nil {
		return response.OrderResponse{}, errors.New("Error while updating order details")
	}
	query1 := `select o.order_id, o.total_amount, o.payment_status, o.order_status, o.delivery_status, o.address_id, p.payment_method from orders as o left join payment_methods as p on o.payment_method_id = p.id where o.order_id=?`
	err1 := o.DB.Raw(query1, uporder.Order_Id).Scan(&orderdetails).Error
	if err1 != nil {
		return response.OrderResponse{}, errors.New("Failed to display order details")
	}
	return orderdetails, nil
}

// List all orders User side
func (o *OrderDatabase) ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	db := o.DB.Model(&domain.Order{})

	query := `select o.order_id,o.total_amount,o.payment_status,o.order_status,o.delivery_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.id where user_id=$1 limit $2 offset $3;`
	if db.Raw(query, userId, limit, offset).Scan(&orders).Error != nil {
		return orders, errors.New("failed to get orders from database")
	}
	return orders, nil
}

// List all orders Admin side
func (o *OrderDatabase) GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	db := o.DB.Model(&domain.Order{})
	query := `select o.order_id,o.total_amount,o.order_status,o.payment_status,o.delivery_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.id limit $1 offset $2;`
	if db.Raw(query, limit, offset).Scan(&orders).Error != nil {
		return orders, errors.New("failed to get orders from database")
	}
	return orders, nil
}

// Delete order
func (o *OrderDatabase) DeleteOrder(c context.Context, order_id uint) error {
	var order domain.Order
	status := "Order cancelled"
	query := `UPDATE orders SET order_status=? WHERE order_id=?`
	err1 := o.DB.Raw(query, status, order_id).Scan(&order).Error
	if err1 != nil {
		return errors.New("failed to update cancel order details")

	}

	err := o.DB.Where("order_id = ?", order_id).Delete(&order).Error
	if err != nil {
		return errors.New("failed to delete order")

	}

	return nil
}

func (o *OrderDatabase) FindOrder(c context.Context, order domain.Order) error {
	status := "Order cancelled"
	query := `select * from orders where order_id=? and order_status=?`
	err := o.DB.Raw(query, order.Order_Id, status).Scan(&order).Error
	fmt.Println("hello", order, "error", err)
	if order.Order_Id == 0 {
		fmt.Println("Order Not Found")
		return errors.New("Order Not Found")
	}
	if order.Order_Status == status {
		return errors.New("Order already cancelled")
	}
	return nil
}

// Place order- Apply coupon
func (o *OrderDatabase) PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error) {
	var paymentresp response.PaymentResponse
	query := `update orders set total_amount=?,  order_status=?, order_date=? where order_id=?`
	err := o.DB.Raw(query, order.Total_Amount, order.Order_Status, order.OrderDate, order.Order_Id).Scan(&order).Error
	if err != nil {
		return response.PaymentResponse{}, errors.New("Failed to update payment")
	}
	query1 := `select order_id, total_amount, order_status, address_id, payment_method_id, payment_status from orders where order_id=?`
	err1 := o.DB.Raw(query1, order.Order_Id).Scan(&paymentresp).Error
	if err1 != nil {
		return response.PaymentResponse{}, errors.New("Failed to display order details")
	}
	return paymentresp, nil
}

// Order status
func (o *OrderDatabase) UpdateOrderStatus(c context.Context, order_id uint, order_status string) (response.OrderResponse, error) {
	var order domain.Order
	var orderResp response.OrderResponse
	query := `update orders set order_status=?  where order_id=?`
	err := o.DB.Raw(query, order_status, order_id).Scan(&order).Error
	if err != nil {
		return response.OrderResponse{}, errors.New("failed to update order status")
	}
	query1 := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.id where o.order_id=?`
	err1 := o.DB.Raw(query1, order_id).Scan(&orderResp).Error
	if err1 != nil {
		return response.OrderResponse{}, errors.New("failed to display order details")
	}
	return orderResp, nil
}

// to place order calculating total amount
func (o *OrderDatabase) FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error) {
	var total_amount float64
	query := `SELECT total_amount FROM orders WHERE order_id=?`
	err := o.DB.Raw(query, order_id).Scan(&total_amount).Error
	if err != nil {
		return 0, errors.New("failed to fetch total amount")
	}
	return total_amount, nil
}

func (o *OrderDatabase) FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error) {
	var order domain.Order
	err := o.DB.Raw("SELECT * FROM orders WHERE order_id=?", order_id).First(&order).Error
	if err != nil {

		return 0, errors.New("failed to find payment method id")
	}
	return uint(order.PaymentMethodID), nil
}
func (o *OrderDatabase) ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error) {
	var returnres response.ReturnResponse
	err := o.DB.Create(&returnOrder).Error
	if err != nil {
		return response.ReturnResponse{}, errors.New("failed to return order , database error")
	}

	query := `select id,order_id,request_date,return_reason,refund_amount,return_status from order_returns where order_id=?`
	err1 := o.DB.Raw(query, returnOrder.OrderID).Scan(&returnres).Error
	if err1 != nil {
		return response.ReturnResponse{}, errors.New("failed to display order details")
	}
	return returnres, nil
}
func (o *OrderDatabase) VerifyOrderID(c context.Context, id uint, orderid uint) error {
	var order domain.Order
	err := o.DB.Where("user_id=? AND order_id=?", id, orderid).First(&order).Error
	if err != nil {
		return errors.New("invalid order id")
	}
	return nil
}

func (o *OrderDatabase) GetAllPendingReturnOrder(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `SELECT 
    r.id AS return_id,
    o.user_id,
    o.order_id,
    r.request_date,
    o.order_date AS requested_at,
    pm.payment_method,
    ps.status AS payment_status,
    r.return_reason AS reason,
    o.total_amount AS order_total,
    r.is_approved
FROM 
    order_returns r
LEFT JOIN 
    orders o ON o.id = r.order_id
LEFT JOIN 
    payment_methods pm ON pm.id = o.payment_method_id
LEFT JOIN 
    payment_details pd ON pd.order_id = o.id
LEFT JOIN 
    payment_statuses ps ON ps.id = pd.payment_status_id
WHERE 
    r.is_approved = false
ORDER BY 
    r.request_date ASC
 LIMIT $1 OFFSET $2`
	err = o.DB.Raw(query, limit, offset).Scan(&ReturnRequests).Error
	if err != nil {
		return ReturnRequests, err
	}
	return ReturnRequests, nil
}
