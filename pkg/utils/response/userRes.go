package response

import (
	"glamgrove/pkg/domain"
	"time"
)

type UserSignUp struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name" copier:"must"`
}
type Profile struct {
	//ID             uint    `json:"id"`
	UserName       string  `json:"user_name" copier:"must"`
	FirstName      string  `json:"first_name" copier:"must"`
	LastName       string  `json:"last_name" copier:"must"`
	Age            uint    `json:"age" copier:"must"`
	Email          string  `json:"email" copier:"must"`
	Phone          string  `json:"phone" copier:"must"`
	DefaultAddress Address `json:"default_address"`
}

type UserResp struct {
	ID          uint      `json:"id" copier:"must"`
	UserName    string    `json:"user_name" copire:"must"`
	FirstName   string    `json:"first_name" copier:"must"`
	LastName    string    `json:"last_name" copier:"must"`
	Age         uint      `json:"age" copier:"must"`
	Email       string    `json:"email" copier:"must"`
	Phone       string    `json:"phone" copier:"must"`
	BlockStatus bool      `json:"block_status" copier:"must"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type Address struct {
	ID        uint   `json:"address_id"`
	House     string `json:"house"`
	City      string `json:"city"`
	State     string `json:"state"`
	PinCode   string `json:"pin_code"`
	Country   string `json:"country"`
	IsDefault bool   `json:"is_default"`
}
type CartItemResp struct {
	CartID   uint   `json:"cart_id"`
	Name     string `json:"product_name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	SubTotal uint   `json:"sub_total"`
}
type CartResp struct {
	TotalProductItems uint    `json:"total_product_items"`
	TotalQty          uint    `json:"total_qty"`
	TotalPrice        float64 `json:"total_price"`
	DiscountAmount    float64 `json:"discount"`
	AppliedCouponID   uint    `json:"applied_coupon_id"`
	AppliedCouponCode string  `json:"applied_coupon_code"`
	CouponDiscount    float64 `json:"coupon_discount"`
	FinalPrice        uint    `json:"final_price"`
	DefaultShipping   Address `json:"default_shipping"`
}
type CheckoutOrder struct {
	UserID         uint           `json:"-"`
	CartItemResp   []CartItemResp `json:"cart_items"`
	TotalQty       uint           `json:"total_qty"`
	TotalPrice     uint           `json:"total_price"`
	Discount       uint           `json:"discount"`
	DefaultAddress domain.Address `json:"address"`
}
type UserContact struct {
	Email string
	Phone string
}
