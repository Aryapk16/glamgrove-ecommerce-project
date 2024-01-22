package usecase

import (
	"context"
	"errors"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"

	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

	"github.com/jinzhu/copier"
)

type adminService struct {
	adminRepository   interfaces.AdminRepository
	PaymentRepository interfaces.PaymentRepository
}

func NewAdminService(repo interfaces.AdminRepository, PaymentRepo interfaces.PaymentRepository) service.AdminService {
	return &adminService{adminRepository: repo,
		PaymentRepository: PaymentRepo}
}

func (a *adminService) Login(c context.Context, admin domain.Admin) (domain.Admin, error) {
	// Check admin exist in db
	// dbAdmin, err := a.adminRepository.GetAdmin(c, admin)
	// if err != nil {
	// 	return admin, err
	// }
	username, password := config.GetAdminDetails()
	dbAdmin := domain.Admin{
		UserName: username,
		Password: password,
	}
	if dbAdmin.UserName != admin.UserName {
		return domain.Admin{}, errors.New("invalid username")
	}
	if dbAdmin.Password != admin.Password {
		return domain.Admin{}, errors.New("invalid password")
	}
	// compare password with hash password
	// if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
	// 	return admin, errors.New("Wrong password")
	// }
	return dbAdmin, nil

}

// List all users in admin side
func (a *adminService) GetAllUser(c context.Context, page request.ReqPagination) (users []response.UserResp, err error) {
	users, err = a.adminRepository.GetAllUser(c, page)

	if err != nil {
		return nil, err
	}

	// if no error then copy users details to an array response struct
	var response []response.UserResp
	copier.Copy(&response, &users)

	return response, nil
}

// to block or unblock a user
func (a *adminService) BlockUnBlockUser(c context.Context, userID uint) error {

	return a.adminRepository.BlockUnBlockUser(c, userID)
}
func (o *adminService) ApproveReturnOrder(c context.Context, data request.ApproveReturnRequest) error {
	// get payment data
	// ID 2 is for status "Paid"
	payment, err := o.PaymentRepository.GetPaymentDataByOrderId(c, data.OrderID)

	if err != nil {
		return err
	}

	data.OrderTotal = payment.OrderTotal
	err = o.adminRepository.ApproveReturnOrder(c, data)
	if err != nil {
		return err
	}
	return nil
}
