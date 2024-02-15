package handler

import (
	"fmt"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponService service.CouponService
}

func NewCouponHandler(CouponUseCase service.CouponService) *CouponHandler {
	return &CouponHandler{
		couponService: CouponUseCase,
	}
}
func (c *CouponHandler) CreateNewCoupon(ctx *gin.Context) {
	var body request.CreateCoupon
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		ctx.JSON(400, response)
		return
	}
	if err := c.couponService.CreateNewCoupon(ctx, body); err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), body)
		ctx.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Coupon created successfully", nil)
	ctx.JSON(200, response)
}

func (c *CouponHandler) ListAllCoupons(ctx *gin.Context) {
	count, err1 := utils.StringToUint(ctx.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(ctx.Query("page_number"))
	if err2 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}
	coupons, err := c.couponService.GetAllCoupons(ctx, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), nil)
		ctx.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "List of coupons", coupons)
	ctx.JSON(200, response)
}

func (c *CouponHandler) MakeCouponInvalid(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "fields provided are in wrong format", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ids := request.Coupon{Coupon: fmt.Sprint(id)}
	if err := c.couponService.MakeCouponInvalid(ctx, ids); err != nil {
		errorRes := response.ErrorResponse(http.StatusBadRequest, "Coupon cannot be made invalid", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorRes)
		return
	}

	response := response.SuccessResponse(200, "Successfully made Coupon as invalid", nil, nil)
	ctx.JSON(http.StatusOK, response)

}
func (c *CouponHandler) ReActivateCoupon(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "fields provided are in wrong format", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ids := request.Coupon{Coupon: fmt.Sprint(id)}
	if err := c.couponService.MakeCouponInvalid(ctx, ids); err != nil {
		errorRes := response.ErrorResponse(http.StatusBadRequest, "Coupon cannot be reactivated", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorRes)
		return
	}

	response := response.SuccessResponse(200, "Successfully made Coupon as valid", nil, nil)
	ctx.JSON(http.StatusOK, response)

}
