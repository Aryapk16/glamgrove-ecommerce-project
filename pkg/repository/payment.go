package repository

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils/request"

	"gorm.io/gorm"
)

type PaymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.PaymentRepository {
	return &PaymentDatabase{DB: db}
}

// Add payment Method
func (pd *PaymentDatabase) AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	err := pd.DB.Create(&payment).Error
	if err != nil {
		return domain.PaymentMethod{}, errors.New("failed to add payment method")
	}
	return payment, nil
}

// Find payment method
func (pd *PaymentDatabase) FindPaymentMethod(c context.Context, payment domain.PaymentMethod) error {
	var payment_methods domain.PaymentMethod

	fmt.Println(payment)
	err := pd.DB.Raw("SELECT * FROM payment_methods WHERE payment_method=$1", payment.PaymentMethod).Scan(&payment_methods).Error
	if err != nil {

		return errors.New("failed to find payment method")
	}

	fmt.Println(payment_methods.ID)

	if payment_methods.ID > 0 {
		return nil
	}
	return errors.New("payment not found")
}

// Find payment method by ID
func (pd *PaymentDatabase) FindPaymentMethodId(c context.Context, method_id uint) (uint, error) {
	var payment_methods domain.PaymentMethod
	err := pd.DB.Raw("SELECT * FROM payment_methods WHERE id=$1", method_id).First(&payment_methods).Error
	if err != nil {

		return 0, errors.New("failed to find payment method")
	}
	return payment_methods.ID, nil
}

// Get payment methods
func (pd *PaymentDatabase) GetPaymentMethods(ctx context.Context, page request.ReqPagination) (payment []domain.PaymentMethod, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `select *from payment_methods limit $1 offset $2;`
	if err := pd.DB.Raw(query, limit, offset).Scan(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

// Delete payment methods
func (p *PaymentDatabase) DeleteMethod(c context.Context, id uint) error {
	var paymentmethod domain.PaymentMethod
	query := `delete from payment_methods where id=?`
	err := p.DB.Raw(query, id).Scan(&paymentmethod).Error
	if err != nil {
		return errors.New("failed to delete payment method")
	}
	return nil

}

// Update payment methods
func (p *PaymentDatabase) UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	query := `update payment_methods set payment_method=? where id=?`
	err := p.DB.Raw(query, payment.PaymentMethod, payment.ID).Scan(&payment).Error
	if err != nil {
		return domain.PaymentMethod{}, errors.New("failed to update payment method details")
	}
	return payment, nil
}

func (p *PaymentDatabase) GetPaymentDataByOrderId(ctx context.Context, orderId uint) (paymentData domain.PaymentDetails, err error) {
	query := `SELECT *
	FROM payment_details 
	WHERE order_id = $1`
	err = p.DB.Raw(query, orderId).Scan(&paymentData).Error
	if err != nil {
		return paymentData, err
	}
	return paymentData, nil

}
