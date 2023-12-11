package usecase

import (
	"context"
	"errors"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"

	"glamgrove/pkg/utils/request"
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
