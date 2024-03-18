package repository

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/api/middleware"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"time"

	"gorm.io/gorm"
)

type adminDatabase struct {
	DB           *gorm.DB
	userDatabase interfaces.UserRepository
}

func NewAdminRepository(db *gorm.DB, userRepo interfaces.UserRepository) interfaces.AdminRepository {
	return &adminDatabase{DB: db,
		userDatabase: userRepo}
}

//	func (a *adminDatabase) GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
//		query := `SELECT * FROM admins WHERE email =? OR user_name =?`
//		if a.DB.Raw(query, admin.Email, admin.UserName).Scan(&admin).Error != nil {
//			return admin, errors.New("Failed to find admin")
//		}
//		return admin, nil
//	}
func (a *adminDatabase) GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserResp, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT * FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err = a.DB.Raw(query, limit, offset).Scan(&users).Error

	return users, err
}

func (a *adminDatabase) BlockUnBlockUser(ctx context.Context, userID uint) error {
	// Check user if valid or not
	var user domain.User
	query := `SELECT * FROM users WHERE id=?`
	a.DB.Raw(query, userID).Scan(user)
	if user.Email == "" {
		// check user email with user ID
		return errors.New("invalid user id user doesn't exist")
	}

	query = `UPDATE users SET block_status = $1 WHERE id = $2`
	if a.DB.Exec(query, !user.BlockStatus, userID).Error != nil {
		return fmt.Errorf("failed to update user block_status to %v", !user.BlockStatus)
	}
	return nil
}

func (a *adminDatabase) ApproveReturnOrder(c context.Context, data request.ApproveReturnRequest) error {
	var order_return domain.OrderReturn
	query := `UPDATE order_returns
	SET is_approved = $1
	WHERE order_id = $2 AND is_approved = false`

	err := a.DB.Exec(query, data.IsApproved, data.OrderID).Error
	if err != nil {
		return err
	}
	query2 := `SELECT * FROM order_returns WHERE order_id=?`
	err2 := a.DB.Raw(query2, data.OrderID).Scan(&order_return).Error
	if err2 != nil {
		return err
	}
	fmt.Println(order_return)
	// if data.IsApproved {

	// 	err := a.userDatabase.CreditUserWallet(c, domain.Wallet{
	// 		UserID:  data.UserID,
	// 		Balance: float64(order_return.RefundAmount),
	// 		Remark:  data.Comment,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	// 	return errors.New("return request denied by admin")
	// }
	return nil
}

//.......................p

func (a *adminDatabase) DashboardUserDetails(c context.Context) (request.DashboardUser, error) {

	var userDetails request.DashboardUser
	err := a.DB.Raw("select count(*) from users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return request.DashboardUser{}, nil
	}

	err = a.DB.Raw("select count(*) from users where block_status = false").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return request.DashboardUser{}, nil
	}

	return userDetails, nil
}

func (a *adminDatabase) DashBoardOrder(c context.Context) (request.DashboardOrder, error) {

	var orderDetails request.DashboardOrder
	err := a.DB.Raw("select count(*) from orders where order_status = 'Order Created' or order_status = 'Order confirmed'").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return request.DashboardOrder{}, nil
	}

	err = a.DB.Raw("select count(*) from orders where order_status = 'order confirmed payment pending'  ").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return request.DashboardOrder{}, nil
	}

	err = a.DB.Raw("select count(*) from orders").Scan(&orderDetails.TotalOrder).Error
	if err != nil {
		return request.DashboardOrder{}, nil
	}

	return orderDetails, nil
}
func (a *adminDatabase) DashBoardProductDetails(c context.Context) (request.DashBoardProduct, error) {

	var productDetails request.DashBoardProduct
	err := a.DB.Raw("select count(*) from products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return request.DashBoardProduct{}, nil
	}

	err = a.DB.Raw("select count(*) from product_items where qty_in_stock = 0").Scan(&productDetails.OutOfStockProduct).Error
	if err != nil {
		return request.DashBoardProduct{}, nil
	}

	return productDetails, nil
}

func (a *adminDatabase) TotalRevenue(c context.Context) (request.DashboardRevenue, error) {

	var revenueDetails request.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := a.DB.Raw("select coalesce(sum(total_amount),0) from orders where payment_status = 'Payment Done' and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return request.DashboardRevenue{}, nil
	}

	startTime, endTime = middleware.GetTimeFromPeriod("month")
	err = a.DB.Raw("select coalesce(sum(total_amount),0) from orders where payment_status = 'Payment Done'  and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return request.DashboardRevenue{}, nil
	}

	startTime, endTime = middleware.GetTimeFromPeriod("year")
	err = a.DB.Raw("select coalesce(sum(total_amount),0) from orders where payment_status = 'Payment Done'  and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return request.DashboardRevenue{}, nil
	}

	return revenueDetails, nil
}

func (a *adminDatabase) AmountDetails(c context.Context) (request.DashboardAmount, error) {

	var amountDetails request.DashboardAmount
	err := a.DB.Raw("select coalesce(sum(total_amount),0) from orders where payment_status = 'Payment Done' or order_status = 'order confirmed' or delivery_status = 'Order delivered successfully'").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return request.DashboardAmount{}, nil
	}

	err = a.DB.Raw("select coalesce(sum(total_amount),0) from orders where payment_status = 'Pending'  or delivery_status = 'Pending' or order_status = 'order confirmed payment pending' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return request.DashboardAmount{}, nil
	}

	return amountDetails, nil

}
func (a *adminDatabase) FilteredSalesReport(c context.Context, startTime time.Time, endTime time.Time) (request.SalesReport, error) {

	var salesReport request.SalesReport

	result := a.DB.Raw("select coalesce(sum(total_amount),0) from orders where payment_status = 'Payment Done'  or created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.TotalSales)
	if result.Error != nil {
		return request.SalesReport{}, result.Error
	}

	result = a.DB.Raw("select count(*) from orders").Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return request.SalesReport{}, result.Error
	}

	result = a.DB.Raw("select count(*) from orders where payment_status = 'Payment Done '  or  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return request.SalesReport{}, result.Error
	}

	result = a.DB.Raw("select count(*) from orders where payment_status = 'Pending'  or  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return request.SalesReport{}, result.Error
	}

	var productID int
	// result = a.DB.Raw("select id from orders group by product_id order by sum(quantity) desc limit 1").Scan(&productID)
	// if result.Error != nil {
	// 	return request.SalesReport{}, result.Error
	// }

	result = a.DB.Raw("select name from products where id = ?", productID).Scan(&salesReport.TrendingProduct)
	if result.Error != nil {
		return request.SalesReport{}, result.Error
	}

	return salesReport, nil
}
