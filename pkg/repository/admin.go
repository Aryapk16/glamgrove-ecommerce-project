package repository

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

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
func (ad *adminDatabase) FindAllUsers(c context.Context, pagination utils.Pagination) ([]response.AllUsers, utils.Metadata, error) {
	var users []response.AllUsers
	var totalrecords int64

	db := ad.DB.Model(&domain.User{})

	//count all records
	if err := db.Count(&totalrecords).Error; err != nil {
		return []response.AllUsers{}, utils.Metadata{}, err
	}

	// Apply pagination
	//db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

	err := db.Raw("select user_id as id,username,name,phone,email from users LIMIT $1 OFFSET $2", pagination.Limit(), pagination.Offset()).Scan(&users).Error
	if err != nil {
		return []response.AllUsers{}, utils.Metadata{}, errors.New("failed to find all users")
	}
	// Compute metadata
	metadata := utils.ComputeMetadata(&totalrecords, &pagination.Page, &pagination.PageSize)

	return users, metadata, nil

}
func (ad *adminDatabase) FindByUsername(c context.Context, Username string) (domain.Admin, error) {
	var admin domain.Admin

	err := ad.DB.Raw("select *from admin_details where username=?", Username).Scan(&admin).Error
	if err != nil {
		return domain.Admin{}, errors.New("failed find user details")
	}
	return admin, nil
}
func (ad *adminDatabase) BlockUser(c context.Context, status request.BlockStatus) error {
	fmt.Println("repohhhhh.........")
	var user domain.User
	ad.DB.Raw("select *from users where user_id=?", status.UserID).Scan(&user)
	if user.ID == 0 {
		return errors.New("user doesn't exist")
	}
	query := `update users set block_status=? where user_id=?`
	err := ad.DB.Raw(query, status.BlockStatus, status.UserID).Scan(&user).Error
	if err != nil {
		return errors.New("failed to update block status")
	}
	return nil
}
