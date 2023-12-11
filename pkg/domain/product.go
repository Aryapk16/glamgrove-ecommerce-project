package domain

// represent a model of product
type Product struct {
	ProductID   uint   `json:"id" gorm:"primaryKey;not null"`
	ProductName string `json:"product_name" gorm:"not null" validate:"required,min=5,max50"`
	Description string `json:"description" gorm:"not null" validate:"required,min=10,max=100"`
	OutOfStock  bool   `json:"out_of_stock" gorm:"not null"`
	Price       uint   `json:"price" gorm:"not null" validate:"required,numeric"`
	QtyInStock  uint   `json:"qtyInStock" gorm:"not null"`
	Image       string `json:"image" gorm:"not null"`
	CategoryID  uint   `json:"categoryId" gorm:"not null"`
	AdminId     uint
	ProductCode  string

	FkCategory Category `json:"-" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}

// for a products category main and sub category as self joining
type Category struct {
	ID           uint   `json:"id" gorm:"primaryKey;not null"`
	CategoryName string `json:"category_name" gorm:"not null"`
}
