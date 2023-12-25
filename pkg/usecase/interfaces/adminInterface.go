package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type AdminUseCase interface {
	AdminLogin(ctx context.Context, admin request.AdminLoginRequest) (domain.Admin, error)
	FindAllUsers(c context.Context, pagination utils.Pagination) ([]response.AllUsers, utils.Metadata, error)

	//AddCategory(ctx context.Context, productCategory domain.Category) (domain.Category, any)

	//FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
	BlockUser(c context.Context, id int) error
	UnBlockUser(c context.Context, id int) error
}
