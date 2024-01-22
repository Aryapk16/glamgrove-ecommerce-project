package response

type PaymentOptionResp struct {
	Id   uint   `json:"code"`
	Name string `json:"option"`
}
type PaymentResponse struct {
	Total_Amount    float64 `json:"total_amount"  gorm:"not null" `
	PaymentMethodId string  `json:"paymentmethod_id"  gorm:"not null" `
	Payment_Status  string  `json:"payment_status"   `
	Order_Status    string  `json:"order_status"`
	Address_Id      uint    `json:"address_id" `
	//Balance_Amount int     `json:"balance_amount"`
	//Payment_Id     string  `json:"payment_id"`

}
