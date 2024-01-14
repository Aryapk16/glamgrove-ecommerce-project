 package request

// // type SignUpReq struct {
// // 	FirstName string `json:"first_name"  validate:"required,min=2,max=50"`
// // 	LastName  string `json:"last_name"  validate:"required,min=1,max=50"`
// // 	Age       uint   `json:"age" validate:"required,numeric"`
// // 	Email     string `json:"email" validate:"required,email"`
// // 	Phone     string `json:"phone" validate:"e164"`
// // 	Password  string `json:"password" validate:"required"`
// // }
// // type LoginRequest struct {
// // 	Email    string `json:"email" binding:"omitempty,email"`
// // 	Password string `json:"Password" binding:"required,min=3,max=30"`
// // }
// // type OtpStruct struct {
// // 	OTP string `json:"otp" validate:"required,min=6,max=6"`
// // }
// // type CodeRequest struct {
// // 	Code string `json:"code"`
// // }
// type UserDetails struct {
// 	Name            string `json:"name"`
// 	Email           string `json:"email" validate:"email"`
// 	Phone           string `json:"phone"`
// 	Password        string `json:"password"`
// 	ConfirmPassword string `json:"confirmpassword"`
// }
// type Address struct {
// Id      uint   `json:"id" gorm:"unique;not null"`
// 	UserID    uint   `json:"user_id"`
// 	Name      string `json:"name" validate:"required"`
// 	HouseName string `json:"house_name" validate:"required"`
// 	Street    string `json:"street" validate:"required"`
// 	City      string `json:"city" validate:"required"`
// 	State     string `json:"state" validate:"required"`
// 	Pin       string `json:"pin" validate:"required"`
// }

// type UserDetailsResponse struct {
// 	Id    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email" validate:"email"`
// 	Phone string `json:"phone"`
// }

// type TokenUsers struct {
// 	Users UserDetailsResponse
// 	Token string
// }

// type UserLogin struct {
// 	Email    string `json:"email" validate:"email"`
// 	Password string `json:"password"`
// }
// type UserSignInResponse struct {
// 	Id       uint   `json:"id"`
// 	UserID   uint   `json:"user_id"`
// 	Name     string `json:"name"`
// 	Email    string `json:"email" validate:"email"`
// 	Phone    string `Json:"phone"`
// 	Password string `json:"password"`
// }