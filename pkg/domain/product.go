package domain

import "gorm.io/gorm"

type Category struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CategoryName string `json:"category_name" gorm:"unique;not null"`
	//Products     []*Product `json:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Product struct
type Product struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Name        string `json:"product_name" gorm:"not null;size:50"`
	Description string `json:"description" gorm:"not null;size:500"`
	CategoryID  uint   `json:"brand_id" gorm:"index;not null"`
	//Category      *Category `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price         uint `json:"price" gorm:"not null"`
	DiscountPrice uint `json:"discount_price" gorm:"default:null"`
}

// ProductItem struct
type ProductItem struct {
	gorm.Model
	ID          uint     `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ProductID   uint     `json:"product_id" gorm:"index;"`
	QtyInStock  uint     `json:"qty_in_stock" gorm:"not null"`
	StockStatus bool     `json:"stock_status" gorm:"not null;default:true;type:boolean;"`
	Price       uint     `json:"price" gorm:"not null"`
	CategoryID  uint     `json:"category_id" gorm:"index;"`
	Category    Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	//SKU           string `json:"sku" gorm:"unique;not null"`
	DiscountPrice uint `json:"discount_price" gorm:"default:null"`
}

type ProductImage struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductId uint   `json:"product_id"`
	Image     string `JSON:"Image" `
}
type ProductItemImage struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	ProductItemID uint   `json:"product_item_id"`
	Image         string `JSON:"Image" `
}
