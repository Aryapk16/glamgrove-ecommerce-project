package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils/request"
)

type PaymentUseCase struct {
	PaymentRepository interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository) service.PaymentService {
	return &PaymentUseCase{PaymentRepository: repo}
}

func (p *PaymentUseCase) AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	fmt.Println("use", payment)
	err := p.PaymentRepository.FindPaymentMethod(c, payment)
	fmt.Println(err)
	if err == nil {
		return domain.PaymentMethod{}, errors.New("Payment method already existsssssssssss")
	}
	paymentresp, err1 := p.PaymentRepository.AddPaymentMethod(c, payment)
	if err1 != nil {
		return domain.PaymentMethod{}, errors.New("Failed to add payment method")
	}
	return paymentresp, nil
}
func (p *PaymentUseCase) GetPaymentMethods(ctx context.Context, page request.ReqPagination) (payment []domain.PaymentMethod, err error) {
	if payment, err = p.PaymentRepository.GetPaymentMethods(ctx, page); err != nil {
		return nil, err
	}
	return payment, nil
}

func (p *PaymentUseCase) DeleteMethod(c context.Context, id uint) error {

	err1 := p.PaymentRepository.DeleteMethod(c, id)
	if err1 != nil {
		return err1
	}
	return nil
}

func (p *PaymentUseCase) UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	//Checking whether the payment id exist
	_, err := p.PaymentRepository.FindPaymentMethodId(c, payment.ID)

	if err != nil {
		return domain.PaymentMethod{}, errors.New("Payment method doesn't exists")
	}

	paymentresp, err := p.PaymentRepository.UpdatePaymentMethod(c, payment)
	if err != nil {
		return domain.PaymentMethod{}, err
	}
	return paymentresp, nil
}
func (p *PaymentUseCase) GetPaymentDataByOrderId(ctx context.Context, orderId uint) (paymentData domain.PaymentDetails, err error) {
	paymentData, err = p.PaymentRepository.GetPaymentDataByOrderId(ctx, orderId)
	if err != nil {
		return paymentData, err
	}
	return paymentData, nil
}
