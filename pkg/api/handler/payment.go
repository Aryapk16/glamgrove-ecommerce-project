package handler

import (
	"glamgrove/pkg/domain"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentService service.PaymentService
}

func NewPaymentHandler(payUseCase service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		PaymentService: payUseCase,
	}
}
func (p *PaymentHandler) AddpaymentMethod(c *gin.Context) {
	var payment domain.PaymentMethod
	if err := c.ShouldBindJSON(&payment); err != nil {
		response := response.ErrorResponse(400, "Error while fetching data from user", err.Error(), payment)
		c.JSON(400, response)
		return
	}

	paymentresp, err1 := p.PaymentService.AddPaymentMethod(c, payment)
	if err1 != nil {
		response := response.ErrorResponse(400, "Can't add payment method", err1.Error(), paymentresp)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully added payment method", paymentresp)
	c.JSON(200, response)
}

func (p *PaymentHandler) GetPaymentMethods(ctx *gin.Context) {
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
	payment, err := p.PaymentService.GetPaymentMethods(ctx, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), nil)
		ctx.JSON(500, response)
		return
	}

	response := response.SuccessResponse(200, "List of payment methods", payment)
	ctx.JSON(200, response)
}

func (p *PaymentHandler) DeleteMethod(c *gin.Context) {
	methodID, err := strconv.Atoi(c.Query("methodID"))

	if err != nil {
		response := response.ErrorResponse(400, "Please add id as params", err.Error(), methodID)
		c.JSON(400, response)
		return
	}

	err1 := p.PaymentService.DeleteMethod(c, uint(methodID))
	if err1 != nil {
		response := response.ErrorResponse(400, "can't delete payment method", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted method")
	c.JSON(200, response)
}

func (p *PaymentHandler) UpdatePaymentMethod(c *gin.Context) {
	var payment domain.PaymentMethod
	if err := c.BindJSON(&payment); err != nil {
		response := response.ErrorResponse(400, "Error while getting data from admin side", err.Error(), payment)
		c.JSON(400, response)
		return
	}
	paymentresp, err := p.PaymentService.UpdatePaymentMethod(c, payment)
	if err != nil {
		response := response.ErrorResponse(400, "Can't update data", err.Error(), "")
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully updated product", paymentresp)
	c.JSON(200, response)
}
