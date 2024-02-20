package interfaces

import (
	"context"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"time"
)

type AdminRepository interface {

	//GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserResp, err error)
	BlockUnBlockUser(ctx context.Context, userID uint) error
	ApproveReturnOrder(c context.Context, data request.ApproveReturnRequest) error
	// .................
	DashboardUserDetails(c context.Context) (request.DashboardUser, error)
	DashBoardOrder(c context.Context) (request.DashboardOrder, error)
	DashBoardProductDetails(c context.Context) (request.DashBoardProduct, error)
	TotalRevenue(c context.Context) (request.DashboardRevenue, error)
	AmountDetails(c context.Context) (request.DashboardAmount, error)

	FilteredSalesReport(ctx context.Context, startTime time.Time, endTime time.Time) (request.SalesReport, error)
}
