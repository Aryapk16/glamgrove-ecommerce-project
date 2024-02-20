package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type AdminService interface {
	Login(c context.Context, admin domain.Admin) (domain.Admin, error)
	GetAllUser(c context.Context, page request.ReqPagination) (users []response.UserResp, err error)
	BlockUnBlockUser(c context.Context, userID uint) error
	ApproveReturnOrder(c context.Context, data request.ApproveReturnRequest) error

	//>.........
	DashBoard(c context.Context) (request.CompleteAdminDashboard, error)
	FilteredSalesReport(c context.Context, timePeriod string) (request.SalesReport, error)
}
