package models

type SignUpReq struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       uint   `json:"age"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
