package repository

import (
	"context"
	"errors"
	"glamgrove/pkg/domain"
	repository "glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils/response"
	"time"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &userDatabase{DB: DB}
}

func (u *userDatabase) FindUser(c context.Context, user domain.User) (domain.User, error) {
	query := `SELECT * FROM users where id=? OR user_name=? OR email=? OR phone=?`
	if err := u.DB.Raw(query, user.ID, user.UserName, user.Email, user.Phone).Scan(&user).Error; err != nil {
		return user, errors.New("Failed to find user")
	}
	return user, nil
}

// Save the user if the user is not existing
func (u *userDatabase) SaveUser(c context.Context, user domain.User) (response.UserSignUp, error) {
	var usersignup response.UserSignUp
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	if u.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age, user.Email, user.Phone, user.Password, createdAt).Error != nil {
		return response.UserSignUp{}, errors.New("Failed to save user")
	}
	query2 := `SELECT id, user_name from users where first_name=?`
	if err := u.DB.Raw(query2, user.FirstName).Scan(&usersignup).Error; err != nil {
		return response.UserSignUp{}, errors.New("Failed to find user")
	}
	return usersignup, nil
}

func (i *userDatabase) GetUserbyID(ctx context.Context, userId uint) (domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE id = ?`
	if err := i.DB.Raw(query, userId).Scan(&user).Error; err != nil {
		return user, err
	}
	//fmt.Println(userId)
	return user, nil
}
