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
func (ad *adminDatabase) FindAdmin(c context.Context, Username string) (domain.AdminDetails, error) {
	var admin domain.AdminDetails
	query := `select * from admins where user_name=?`
	err := ad.DB.Raw(query, Username).Scan(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("admin not found")
	}

	return admin, nil
}

// for adding admin to database
func (ad *adminDatabase) AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Create(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while adding admin details to database")
	}

	return admin, nil
}
func (ad *adminDatabase) FindByUsername(c context.Context, Username string) (domain.AdminDetails, error) {
	var admin domain.AdminDetails

	err := ad.DB.Raw("select *from admin_details where username=?", Username).Scan(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("failed find user details")
	}
	return admin, nil
}
