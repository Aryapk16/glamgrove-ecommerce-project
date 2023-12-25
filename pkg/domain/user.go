package domain

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"`
	FirstName   string `json:"first_name" gorm:"not null" validate:"required,min=2,max=50"`
	LastName    string `json:"last_name" gorm:"not null" validate:"required,min=1,max=50"`
	Age         uint   `json:"age" gorm:"not null" validate:"required,numeric"`
	Email       string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Phone       string `json:"phone" gorm:"unique;not null" validate:"required,min=10,max=15"`
	Password    string `json:"password" gorm:"not null" validate:"required"`
	Status      string `json:"status" gorm:"not null"`
	BlockStatus bool   `json:"block_status" gorm:"not null;default:false"`
}
