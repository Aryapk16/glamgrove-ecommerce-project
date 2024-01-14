package handler

import (
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
}

func NewAdminHandler(adminService interfaces.AdminService) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminService,
	}
}
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
