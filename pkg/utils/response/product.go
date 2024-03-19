package response

import "time"

type ResponseProduct struct {
	ID            uint   `json:"id"`
	Name          string `json:"product_name"`
	Description   string `json:"description" `
	Category_name string `json:"category_name"`
	//Price         uint   `json:"price"`
	//DiscountPrice uint   `json:"discount_price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductItemResp struct {
	ProductID      uint   `json:"product_id"`
	ProductItemID  uint   `json:"product_item_id"`
	ProductBrand   string `json:"product_brand"`
	StockAvailable uint   `json:"stock_available"`
	ProductName    string `json:"product_name"`
	Category       string `json:"brand"`
	Description    string `json:"description"`

	Price      uint     `json:"price"`
	OfferPrice uint     `json:"offer_price"`
	Images     []string `json:"images"`
}

type Brand struct {
	ID           uint   `json:"Brand_id"`
	CategoryName string `json:"Brand_name"`
}
type SalesResponse struct {
	TotalOrder     int64 `json:"total_order"`
	DeliveredOrder int64 `json:"delivered_order"`
	PendingOrder   int64 `json:"pending_order"`
	CancelledOrder int64 `json:"canncelled_order"`
}
