package request

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string `json:"password" validate:"required,min=8,max=64" `
}
type Block struct {
	UserID uint `json:"user_id" binding:"required,numeric"`
}

type BlockStatus struct {
	UserID      uint `json:"user_id" binding:"required,numeric"`
	BlockStatus bool `json:"blockstatus"`
}
type DeleteId struct {
	ProductID uint `json:"productid" binding:"required,numeric"`
}
