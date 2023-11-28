package request

type SignUpReq struct {
	FirstName string `json:"first_name" gorm:"not null" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" gorm:"not null" validate:"required,min=1,max=50"`
	Age       uint   `json:"age" gorm:"not null" validate:"required,numeric"`
	Email     string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Phone     string `json:"phone" gorm:"unique;not null" validate:"required,min=10,max=10"`
	Password  string `json:"password" gorm:"not null" validate:"required"`
}
type LoginData struct {
	UserName string `json:"user_name" binding:"omitempty,min=3,max=15"`
	//Phone    string `json:"phone" binding:"omitempty,min=10,max=10"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"Password" binding:"required,min=3,max=30"`
}
