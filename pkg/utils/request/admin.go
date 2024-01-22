package request

type AdminLogin struct {
	UserName string `json:"user_name" validate:"min=8,max=20"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type ApproveReturnRequest struct {
	ReturnID   uint   `json:"return_id"`
	OrderID    uint   `json:"order_id"`
	UserID     uint   `json:"user_id"`
	OrderTotal uint   `json:"-"`
	IsApproved bool   `json:"is_approved"`
	Comment    string `json:"comment"`
}
