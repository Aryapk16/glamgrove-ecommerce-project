package interfaces

import (
	"context"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type AdminRepository interface {

	//GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserResp, err error)
	BlockUnBlockUser(ctx context.Context, userID uint) error
}
