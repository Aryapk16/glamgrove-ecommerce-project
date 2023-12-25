package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type AdminRepository interface {
	FindAdmin(c context.Context, Username string) (domain.Admin, error)
	AddAdmin(c context.Context, admin domain.Admin) (domain.Admin, error)

	//user management
	BlockUser(c context.Context, status request.BlockStatus) error
	FindByUsername(c context.Context, Username string) (domain.Admin, error)
	FindAllUsers(c context.Context, pagination utils.Pagination) ([]response.AllUsers, utils.Metadata, error)
	//UnBlockUser(c context.Context, id int) error
}
