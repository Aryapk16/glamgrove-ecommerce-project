package request

type UpdateOrder struct {
	Order_Id        uint   ` json:"order_id"`
	PaymentMethodID uint   `json:"paymentmethod_id"  gorm:"not null" `
	Address_Id      uint   `json:"address_id" `
	Payment_Status  string `json:"payment_status"`
	DeliveryStatus  string `json:"delivery_status"`
}

type OrderRequest struct {
	Address_Id      int `json:"address_id" binding:"required,numeric"`
	PaymentMethodID int `json:"paymentmethod_id"  gorm:"not null" `
}
type PlaceOrderRequest struct {
	OrderId  int `json:"order_id"`
	CouponId int `json:"coupon_id,omitempty"`
}

type RazorpayVeification struct {
	RazorpayPaymentID string `json:"razorpay_payment_id"`
	RazorpayOrderID   string `json:"razorpay_order_id"`
	RazorpaySignature string `json:"razorpay_signature"`
	UserID            string `json:"user_id"`
}
type RazorPayReq struct {
	Total_Amount float64
	Email        string
	Phone        string
}

type ReturnRequest struct {
	UserID  uint   `json:"-"`
	OrderID uint   `json:"order_id"`
	Reason  string `json:"reason"  binding:"omitempty,min=4,max=15"`
}
