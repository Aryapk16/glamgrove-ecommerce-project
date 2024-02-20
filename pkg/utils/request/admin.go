package request

import "time"

type AdminLogin struct {
	UserName string `json:"user_name" validate:"min=8,max=20"`
	Password string `json:"password" validate:"min=8,max=20"`
}
type ReqSalesReport struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	//Pagination utils.Pagination `json:"pagination"`
}

type ApproveReturnRequest struct {
	ReturnID   uint   `json:"return_id"`
	OrderID    uint   `json:"order_id"`
	UserID     uint   `json:"user_id"`
	OrderTotal uint   `json:"-"`
	IsApproved bool   `json:"is_approved"`
	Comment    string `json:"comment"`
}

type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	TrendingProduct string
}

// .......................
type DashboardUser struct {
	TotalUsers  int
	BlockedUser int
}

type DashBoardProduct struct {
	TotalProducts     int
	OutOfStockProduct int
}
type DashboardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}
type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}
type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}
type CompleteAdminDashboard struct {
	DashboardUser    DashboardUser
	DashBoardProduct DashBoardProduct
	DashboardOrder   DashboardOrder
	DashboardRevenue DashboardRevenue
	DashboardAmount  DashboardAmount
}
