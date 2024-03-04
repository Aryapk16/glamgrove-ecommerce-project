package handler

import (
	"fmt"
	"glamgrove/pkg/auth"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminHandler struct {
	adminUseCase interfaces.AdminService
	orderService interfaces.OrderService
}

func NewAdminHandler(adminService interfaces.AdminService, orderservice interfaces.OrderService) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminService,
		orderService: orderservice,
	}
}

// AdminLogin godoc
// @Summary Admin login
// @Description Login to Admin account.
// @Tags Admin
// @Accept json
// @Produce json
// @Param  input body request.AdminLogin true "inputs"
// @Success  200 {object} response.Response
// @Failure  400 {object} response.Response
// @Failure  500 {object} response.Response
// @Router /admin/login/ [post]
func (a *AdminHandler) AdminLogin(c *gin.Context) {
	//Bind login data
	var body request.AdminLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
	}
	// validate login data
	var admin domain.Admin
	copier.Copy(&admin, body)
	admin, err := a.adminUseCase.Login(c, admin)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to login", err.Error(), admin)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Setup JWT
	if !auth.JwtCookieSetup(c, "admin-auth", admin.ID) {
		response := response.ErrorResponse(500, "Generate JWT failure", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
	}
	// Success response
	response := response.SuccessResponse(200, "Successfully logged in", nil)
	c.JSON(http.StatusOK, response)
}

// ListUsers godoc
// @Summary Get a list of users
// @Description Get a paginated list of users.
// @Tags Users
// @Accept json
// @Produce json
// @Param  count query int false "Number of users to fetch per page"
// @Param  page_number query int false "Page number"
// @Success  200 {object}	response.Response{} "List user successful"
// @Failure  400 {object} response.Response{} "Missing or invalid inputs"
// @Failure  500 {object} response.Response{} "Failed to get all users"
// @Router /admin/users/ [get]
func (a *AdminHandler) ListUsers(c *gin.Context) {

	count, err1 := utils.StringToUint(c.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))

	if err2 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	users, err := a.adminUseCase.GetAllUser(c, pagination)
	if err != nil {
		respone := response.ErrorResponse(500, "Failed to get all users", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, respone)
		return
	}

	// check there is no user
	if len(users) == 0 {
		response := response.SuccessResponse(200, "Oops!...No user to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := response.SuccessResponse(200, "List user successful", users)
	c.JSON(http.StatusOK, response)

}

// BlockUnBlockUser godoc
// @Summary Block or unblock user
// @Description Blocks or unblocks in admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Param input body request.Block true "inputs"
// @Success         200 {object}    response.Response{}  "Successfully changed user block_status"
//
//	@Failure		400	{object}	response.Response{}		"Invalid inputs"
//	@Failure		400	{object}	response.Response{}		"Failed to change user block_status"
//
// @Router /admin/users/block [post]
func (a *AdminHandler) BlockUnBlockUser(ctx *gin.Context) {
	var body request.Block
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// if successfully blocked or unblock user then response 200
	err := a.adminUseCase.BlockUnBlockUser(ctx, body.UserID)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to change user block_status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully changed user block_status", body.UserID)
	ctx.JSON(http.StatusOK, response)
}

// GetAllReturnOrder godoc
// @Summary Get all pending return requests
// @Description Retrieves a list of all pending return requests with pagination support.
// @Tags Return Orders
// @Accept json
// @Produce json
// @Param count query integer false "Number of items per page"
// @Param page_number query integer false "Page number"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{}	 "Return Request List"
// @Failure 400 {object} response.Response{}	 "Missing or invalid inputs" or "Something went wrong!"
// @Router  /admin/users/return-orders [get]
func (a *AdminHandler) GetAllReturnOrder(c *gin.Context) {

	count, err1 := utils.StringToUint(c.Query("count"))
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))

	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err2 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	fmt.Println(pagination)
	returnRequests, err := a.orderService.GetAllPendingReturnRequest(c, pagination)

	if err != nil {
		response := response.ErrorResponse(400, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Return Request List", returnRequests)
	c.JSON(http.StatusOK, response)

}

func (a *AdminHandler) ApproveReturnOrder(c *gin.Context) {
	var body request.ApproveReturnRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Invalid Request Body", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	body.IsApproved = true
	err := a.adminUseCase.ApproveReturnOrder(c, body)
	if err != nil {
		response := response.ErrorResponse(400, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(body)
	response := response.SuccessResponse(http.StatusOK, "Return Order Approved", body)
	c.JSON(http.StatusOK, response)
}

// ...............................dashboard
// DashBoard godoc
// @Summary Get dashboard details
// @Description  Get details for the dashboard
// @Tags Admin Dashboard
// @Accept json
// @Produce json
// @Success 200 {object}  response.Response{} "succesfully recevied all records"
// @Failure 400	{object}	response.Response{}		"error in getting dashboard details"
// @Router /admin/dashboard/ [get]
func (a *AdminHandler) DashBoard(c *gin.Context) {
	dashBoard, err := a.adminUseCase.DashBoard(c)
	if err != nil {
		errRes := response.ErrorResponse(http.StatusBadRequest, "error in getting dashboard details", err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	sucessRes := response.SuccessResponse(http.StatusOK, "succesfully recevied all records", dashBoard, nil)
	c.JSON(http.StatusOK, sucessRes)
}

// .............
// FilteredSalesReport godoc
// @Summary Get filtered sales report
// @Description Get  sales report for a specific time period
// @Tags Admin
// @Accept json
// @Produce json
// @Param period path string true "Time period for the sales report"
// @Success 200 {object} response.Response{}  "Sales report retrieved successfully"
// @Failure 500 {object} response.Response{}  "Sales report could not be retrieved"
// @Router /admin/dashboard/filteredSalesReport [get]
func (a *AdminHandler) FilteredSalesReport(c *gin.Context) {

	timePeriod := c.Param("period")
	salesReport, err := a.adminUseCase.FilteredSalesReport(c, timePeriod)
	if err != nil {
		errorRes := response.ErrorResponse(http.StatusInternalServerError, "sales report could not be retrieved", err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.SuccessResponse(http.StatusOK, "sales report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, successRes)

}
