package response

import "time"

type OrderResponse struct {
	Order_id uint `json:"order_id"`
	//UserName       uint    `json:"user_name"`
	Total_Amount   float64 `json:"total_amount"  gorm:"not null" `
	Order_Status   string  `json:"order_status"`
	Payment_Status string  `json:"payment_status"   `
	DeliveryStatus string  `json:"delivery_status"`
	Address_Id     uint    `json:"address_id" `
	Payment_Method string  `json:"payment_method"`
}
type PhoneEmailResp struct {
	Phone string
	Email string
}
type ResRazorpayOrder struct {
	RazorpayKey     string      `json:"razorpay_key"`
	UserID          uint        `json:"user_id"`
	AmountToPay     uint        `json:"amount_to_pay"`
	RazorpayOrderID interface{} `json:"razorpay_order_id"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
}

type ReturnResponse struct {
	ID           uint      `gorm:"serial primaryKey;autoIncrement:true;unique"`
	OrderID      uint      `json:"order_id" gorm:"not null;unique"`
	RequestDate  time.Time `json:"request_date" gorm:"not null"`
	ReturnReason string    `json:"return_reason" gorm:"not null"`
	RefundAmount float64   `json:"refund_amount" gorm:"not null"`
	ReturnStatus string    `json:"return_status"`
}

type ReturnRequests struct {
	ReturnID      uint      `json:"return_id"`
	UserID        uint      `json:"user_id"`
	OrderId       uint      `json:"order_id"`
	RequestedAt   time.Time `json:"requested_at"`
	OrderDate     time.Time `json:"order_date"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	Reason        string    `json:"reason"`
	OrderTotal    uint      `json:"order_total"`
	IsApproved    bool      `json:"is_approved"`
}
