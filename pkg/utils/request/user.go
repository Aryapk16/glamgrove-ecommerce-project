package request

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Address struct {
	ID        uint      `json:"-"`
	UserID    uint      `json:"user_id"`
	House     string    `json:"house"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	PinCode   string    `json:"pin_code"`
	Country   string    `json:"country"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
type AddressPatch struct {
	ID        uint      `json:"address_id"`
	UserID    uint      `json:"-"`
	House     string    `json:"house"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	PinCode   string    `json:"pin_code"`
	Country   string    `json:"country"`
	IsDefault bool      `json:"is_default"`
	UpdatedAt time.Time `json:"-"`
}
type AddToCartReq struct {
	UserID         uint    `json:"user_id"`
	ProductID      uint    `json:"product_id" binding:"required"`
	ProductItemID  uint    `json:"product_item_id" binding:"required"`
	Quantity       uint    `json:"quantity" binding:"required"`
	Price          float64 `json:"-"`
	Discount_price uint    `json:"-"`
}
type UpdateCartReq struct {
	UserID        uint `json:"-"`
	ProductID     uint `json:"product_id" binding:"required"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
	Quantity      uint `json:"quantity" binding:"required"`
}
type DeleteCartItem struct {
	UserID        uint `json:"-"`
	ProductID     uint `json:"product_id" binding:"required"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
}

type OtpCookieStruct struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	jwt.StandardClaims
}
