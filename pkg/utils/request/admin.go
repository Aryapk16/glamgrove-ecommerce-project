package request

type AdminLogin struct {
	UserName string `json:"user_name" validate:"min=8,max=20"`
	Password string `json:"password" validate:"min=8,max=20"`
}
