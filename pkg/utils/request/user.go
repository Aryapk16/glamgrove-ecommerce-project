package request

type SignUpReq struct {
	FirstName string `json:"first_name"  validate:"required,min=2,max=50"`
	LastName  string `json:"last_name"  validate:"required,min=1,max=50"`
	Age       uint   `json:"age" validate:"required,numeric"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"e164"`
	Password  string `json:"password" validate:"required"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"Password" binding:"required,min=3,max=30"`
}
type OtpStruct struct {
	OTP string `json:"otp" validate:"required,min=6,max=6"`
}
type CodeRequest struct {
	Code string `json:"code"`
}
