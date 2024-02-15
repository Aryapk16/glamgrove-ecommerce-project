package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
)

type CouponRepository interface {
	CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error
	GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error)
	MakeCouponInvalid(ctx context.Context, id request.Coupon) error
	IsCouponExist(ctx context.Context, id request.Coupon) (error, bool)
	ReActivateCoupon(ctx context.Context, id request.Coupon) error
	GetCouponBycode(ctx context.Context, code string) (coupon domain.Coupon, err error)
	GetCouponById(ctx context.Context, couponId uint) (coupon domain.Coupon, err error)
	ApplyCoupon(ctx context.Context, data utils.ApplyCoupon) (AppliedCoupon utils.ApplyCouponResponse, err error)
}
