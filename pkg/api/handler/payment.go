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

// AddpaymentMethod godoc
// @Summary Add a new payment method
// @Description Adds a new payment method based on the provided details.
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param payment body domain.PaymentMethod true "Payment method details"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully added payment method"
// @Failure 400 {object} response.Response{} "Error while fetching data from user" or "Can't add payment method"
// @Router /admin/paymentmethod/add  [post]
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

// GetPaymentMethods godoc
// @Summary Get payment methods
// @Description Retrieves a list of payment methods with pagination support.
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param count query integer false "Number of items per page"
// @Param page_number query integer false "Page number"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "List of payment methods"
// @Failure 400 {object} response.Response{} "Invalid inputs"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router   /admin/paymentmethod/view [get]
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

// DeleteMethod godoc
// @Summary Delete a payment method
// @Description Deletes a payment method by its ID.
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param methodID query integer true "Payment method ID to delete"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Success message"
// @Failure 400 {object} response.Response{} "Invalid parameters"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /admin/paymentmethod/delete [delete]
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

// UpdatePaymentMethod godoc
// @Summary Update a payment method
// @Description Updates an existing payment method.
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param body body domain.PaymentMethod true "Payment method object to update"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{}  "Success message"
// @Failure 400 {object} response.Response{} "Error while getting data or invalid parameters"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router   /admin/paymentmethod/update [put]
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
