package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserName    string `json:"user_name" gorm:"not null" binding:"required,min=3,max=15"`
	FirstName   string `json:"first_name" gorm:"not null" binding:"required,min=2,max=20"`
	LastName    string `json:"last_name" gorm:"not null" binding:"required,min=2,max=20"`
	Age         uint   `json:"age" gorm:"not null" binding:"required,numeric"`
	Email       string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Phone       string `json:"phone" gorm:"unique;not null" binding:"required,min=10,max=10"`
	Password    string `json:"password" gorm:"not null" binding:"required"`
	BlockStatus bool   `json:"block_status" gorm:"not null;default:false"`
}
type Address struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserID    uint   `json:"-"`
	House     string `json:"house" gorm:"not null"`
	City      string `json:"city" gorm:"not null" binding:"required,min=2,max=20"`
	State     string `json:"state" gorm:"not null" binding:"required,min=2,max=20"`
	PinCode   string `json:"pin_code" gorm:"not null" binding:"required,min=2,max=10"`
	Country   string `json:"country" gorm:"not null" binding:"required,min=2,max=20"`
	IsDefault bool   `gorm:"not null"`
}
type CartItems struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	CartID      uint    `gorm:"not null"`
	ProductID   uint    `gorm:"not null"`
	Quantity    uint    `gorm:"not null"`
	StockStatus bool    `gorm:"not null;default:true"`
	Price       float64 `gorm:"not null"`
}

type Cart struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"not null"`
	Total  float64 `gorm:"default:0"`
}
