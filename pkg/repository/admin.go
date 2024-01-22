package repository

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

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
	a.DB.Raw(query, userID).Scan(&user)
	if user.Email == "" {
		// check user email with user ID
		return errors.New("Invalid user id user doesn't exist")
	}

	query = `UPDATE users SET block_status = $1 WHERE id = $2`
	if a.DB.Exec(query, !user.BlockStatus, userID).Error != nil {
		return fmt.Errorf("Failed to update user block_status to %v", !user.BlockStatus)
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
