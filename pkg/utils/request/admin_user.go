package request

type SignupUserData struct {
UserName  string `json:"user_name"  binding:"required,min=3,max=15"`
FirstName string `json:"first_name"  binding:"required,min=2,max=50"`
LastName  string `json:"last_name"  binding:"required,min=1,max=50"`
Age       uint   `json:"age"  binding:"required,numeric"`
Email     string `json:"email" binding:"required,email"`
Phone     string `json:"phone" binding:"required,min=10,max=10"`
Password  string `json:"password"  binding:"required"`
}
type LoginData struct {
	UserName string `json:"user_name" binding:"omitempty,min=3,max=15"`
	//Phone    string `json:"phone" binding:"omitempty,min=10,max=10"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"Password" binding:"required,min=3,max=30"`
}
type OTPVerify struct {
	OTP    string `json:"otp" binding:"required,min=4,max=8"`
	UserID uint   `json:"user_id" binding:"required,numeric"`
}
type ReqPagination struct {
	Count      uint `json:"count"`
	PageNumber uint `json:"page_number"`
}

type Block struct {
	UserID uint `json:"user_id" binding:"required,numeric"`
}
type UserID struct {
	UserID uint `json:"user_id" binding:"required,numeric"`
}
