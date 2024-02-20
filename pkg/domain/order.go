package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_Id uint ` gorm:" serial primaryKey;autoIncrement:true;unique"`
	User_Id  uint `json:"user_id"  gorm:"not null" `
	//Applied_Coupon_id uint      `json:"applied_coupon_id,omitempty"`
	Total_Amount    float64   `json:"total_amount"  gorm:"not null" `
	PaymentMethodID int       `json:"paymentmethod_id"  gorm:"not null" `
	Payment_Status  string    `json:"payment_status"`
	Order_Status    string    `json:"order_status"`
	DeliveryStatus  string    `json:"delivery_status"`
	Address_Id      int       `json:"address_id" `
	OrderDate       time.Time `json:"order_date"`
}
type OrderReturn struct {
	//ID          uint      `gorm:"serial primaryKey;autoIncrement:true;unique"`
	OrderID      uint      `json:"order_id" gorm:"not null;unique"`
	RequestDate  time.Time `json:"request_date" gorm:"not null"`
	ReturnReason string    `json:"return_reason" gorm:"not null"`
	RefundAmount float64   `json:"refund_amount" gorm:"not null"`
	IsApproved   bool      `json:"is_approved"`
	ReturnDate   time.Time `json:"return_date"`
	ReturnStatus string    `json:"return_status"`
}

//...........................................

type Sale struct {
	UserID          int
	Name            string
	Email           string
	OrderDate       time.Time
	OrderTotalPrice float64
	OrderStatus     string
	PaymentStatus   bool
	// Other fields
}
