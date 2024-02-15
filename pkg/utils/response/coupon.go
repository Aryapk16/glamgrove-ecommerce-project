package response

import "time"

type CouponResponse struct {
	DiscountPercent float64   `json:"discount_percent,omitempty"`
	ValidTill       time.Time `json:"valid_till"`
}
