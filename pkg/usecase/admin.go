package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"

	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type AdminUsecase struct {
	adminRepo interfaces.AdminRepository
}

func NewadminUseCase(repo interfaces.AdminRepository) service.AdminUseCase {
	return &AdminUsecase{adminRepo: repo}
}

func (ad *AdminUsecase) AdminLogin(ctx context.Context, admin request.AdminLoginRequest) (domain.Admin, error) {
	dbAdmin, _ := ad.adminRepo.FindAdmin(ctx, admin.Username)

	// check password matching

	if dbAdmin.Password != admin.Password {
		return domain.Admin{}, errors.New("password is not correct")
	}

	return dbAdmin, nil
}
func (ad *AdminUsecase) FindAllUsers(ctx context.Context, pagination utils.Pagination) ([]response.AllUsers, utils.Metadata, error) {
	users, metadata, err := ad.adminRepo.FindAllUsers(ctx, pagination)
	if err != nil {
		return []response.AllUsers{}, utils.Metadata{}, errors.New("error while finding all users")
	}
	return users, metadata, nil
}
func (ad *AdminUsecase) BlockUser(ctx context.Context, id int) error {
	fmt.Println("usecase.........")
	var status request.BlockStatus
	status.UserID = uint(id)
	status.BlockStatus = true
	err := ad.adminRepo.BlockUser(ctx, status)
	if err != nil {
		return err
	}
	return nil
}
func (ad *AdminUsecase) UnBlockUser(ctx context.Context, id int) error {
	var status request.BlockStatus
	status.UserID = uint(id)
	status.BlockStatus = false
	err := ad.adminRepo.BlockUser(ctx, status)
	if err != nil {
		return err
	}
	return nil
}
