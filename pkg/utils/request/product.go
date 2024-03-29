package request

type ProductReq struct {
	Name        string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Description string `json:"description" gorm:"not null" binding:"required,min=10,max=1000"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	//Image       string `json:"image" gorm:"not null" binding:"required"`
}

type UpdateProductReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"product_name,omitempty"`
	Description string `json:"description,omitempty"`
	CategoryID  uint   `json:"brand_id,omitempty"`
	//Price       uint   `json:"price,omitempty"`
	//Image       string `json:"image,omitempty"`
}
type UpdateProductItemReq struct {
	ID            uint   `json:"id"`
	ProductId     uint   `json:"product_id,omitempty"`
	QtyInStock    uint   `json:"quantity_inStock,omitempty"`
	SKU           string `json:"sku,omitempty"`
	DiscountPrice string `json:"discount_price,omitempty"`
	Price         uint   `json:"price,omitempty"`
	Image         string `json:"image,omitempty"`
}

type DeleteProductReq struct {
	ID uint `json:"Prod_id" binding:"required"`
}

type Category struct {
	// ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}

type ProductItemReq struct {
	ProductID    uint   `json:"product_id" binding:"required"`
	ProductBrand string `json:"product_brand" binding:"required"`

	QtyInStock    uint `json:"qty_in_stock" binding:"required"`
	Price         uint `json:"price"`
	DiscountPrice uint `json:"discount_price,omitempty"`
	//SKU           string `json:"SKU" binding:"required"`
	//Configurations map[string]Variation `json:"configurations" binding:"required"`
	//Images []string `json:"images" binding:"required"`
}
