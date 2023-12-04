package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
)

type AdminUsecase interface {
	AdminLogin(ctx context.Context, admin request.AdminLoginRequest) (domain.AdminDetails, error)
	// FindAllUser(ctx context.Context) ([]domain.Users, error)

	//AddCategory(ctx context.Context, productCategory domain.Category) (domain.Category, any)

	// FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
}
