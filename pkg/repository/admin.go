package repository

import (
	"context"
	"errors"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewadminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}
func (ad *adminDatabase) FindAdmin(c context.Context, Username string) (domain.Admin, error) {
	var admin domain.Admin
	query := `select * from admins where user_name=?`
	err := ad.DB.Raw(query, Username).Scan(&admin).Error
	if err != nil {
		return domain.Admin{}, errors.New("admin not found")
	}

	return admin, nil
}

// for adding admin to database
func (ad *adminDatabase) AddAdmin(c context.Context, admin domain.Admin) (domain.Admin, error) {
	err := ad.DB.Create(&admin).Error
	if err != nil {
		return domain.Admin{}, errors.New("error while adding admin details to database")
	}

	return admin, nil
}
func (ad *adminDatabase) FindByUsername(c context.Context, Username string) (domain.Admin, error) {
	var admin domain.Admin

	err := ad.DB.Raw("select *from admin_details where username=?", Username).Scan(&admin).Error
	if err != nil {
		return domain.Admin{}, errors.New("failed find user details")
	}
	return admin, nil
}
