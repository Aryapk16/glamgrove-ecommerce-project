package response

type AllUsers struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=12"`
	Name     string `json:"name" validate:"required,min=3,max=12" `
	Phone    string `json:"phone" gorm:"unique" binding:"required,min=10,max=10"`
	Email    string `json:"email" validate:"required,min=3,max=12" `
}
// type AdminDetailsResponse struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email" validate:"email"`
// }

// type TockenAdmin struct {
// 	Admin       AdminDetailsResponse
// 	AccessToken string
// 	// RefreshToken string
// }
