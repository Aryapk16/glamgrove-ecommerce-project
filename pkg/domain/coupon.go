package domain

import "time"

type Coupon struct {
	ID                uint   `gorm:"primaryKey" `
	Code              string `gorm:"unique" `
	MinOrderValue     float64
	DiscountPercent   float64
	DiscountMaxAmount float64
	ValidTill         time.Time
	Valid             bool `gorm:"default :true"`
}
