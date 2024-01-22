package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
)

type PaymentRepository interface {
	AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	FindPaymentMethod(c context.Context, payment domain.PaymentMethod) error
	GetPaymentMethods(ctx context.Context, page request.ReqPagination) (payment []domain.PaymentMethod, err error)
	FindPaymentMethodId(c context.Context, method_id uint) (uint, error)
	UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	DeleteMethod(c context.Context, id uint) error
	GetPaymentDataByOrderId(ctx context.Context, orderId uint) (paymentData domain.PaymentDetails, err error)
}
