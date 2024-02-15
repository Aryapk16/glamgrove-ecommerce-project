package usecase

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils/request"
)

type couponUseCase struct {
	couponRepository interfaces.CouponRepository
}

func NewCouponUseCase(CouponRepo interfaces.CouponRepository) service.CouponService {
	return &couponUseCase{couponRepository: CouponRepo}
}
func (c *couponUseCase) CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error {
	if err := c.couponRepository.CreateNewCoupon(ctx, CouponData); err != nil {
		return err
	}
	return nil
}
func (c *couponUseCase) GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error) {
	if coupon, err = c.couponRepository.GetAllCoupons(ctx, page); err != nil {
		return nil, err
	}
	return coupon, nil
}
func (c *couponUseCase) MakeCouponInvalid(ctx context.Context, id request.Coupon) error {
	if err := c.couponRepository.MakeCouponInvalid(ctx, id); err != nil {
		return err
	}
	return nil
}
func (c *couponUseCase) ReActivateCoupon(ctx context.Context, id request.Coupon) error {
	if err := c.couponRepository.ReActivateCoupon(ctx, id); err != nil {
		return err
	}
	return nil
}
