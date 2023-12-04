package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"

	"glamgrove/pkg/utils/request"
)

type AdminUsecase struct {
	adminRepo interfaces.AdminRepository
}

func NewadminUsecase(repo interfaces.AdminRepository) service.AdminUsecase {
	return &AdminUsecase{adminRepo: repo}
}

func (ad *AdminUsecase) AdminLogin(ctx context.Context, admin request.AdminLoginRequest) (domain.AdminDetails, error) {
	dbAdmin, _ := ad.adminRepo.FindAdmin(ctx, admin.Username)

	// check password matching

	// if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
	// 	return domain.AdminDetails{}, errors.New("password is not correct")
	// }

	if dbAdmin.Password != admin.Password {
		return domain.AdminDetails{}, errors.New("password is not correct")
	}

	fmt.Println("hii")
	return dbAdmin, nil
}
