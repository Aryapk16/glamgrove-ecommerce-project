package handler

import (
	"errors"
	"fmt"
	"glamgrove/pkg/auth"
	services "glamgrove/pkg/usecase/interfaces"
	"strconv"

	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUsecase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUsecase: usecase}
}

func (ad *AdminHandler) AdminLogin(ctx *gin.Context) {

	var body request.AdminLoginRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), body)
		utils.ResponseJSON(ctx, res)
		return
	}

	if body.Username == " " && body.Password == " " {
		err := errors.New("enter username and password")
		res := response.ErrorResponse(400, "invalid input", err.Error(), nil)
		utils.ResponseJSON(ctx, res)
		return
	}
	adminData, err := ad.AdminUsecase.AdminLogin(ctx, body)
	if err != nil {
		response := response.ErrorResponse(400, "Enter valid username", err.Error(), "login failed")

		ctx.JSON(400, response)
		return
	}

	if _, err := ad.AdminUsecase.AdminLogin(ctx, body); err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), "login failed")

		ctx.JSON(400, response)
		return
	}

	message := "Successfully logged in as " + body.Username

	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWT(int(adminData.ID))
	if err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), "login failed")

		ctx.JSON(400, response)
		return
	}
	//set cookie

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Admin_Authorization", tokenString["accessToken"], 3600*24*30, "/", " ", false, true)

	response := response.SuccessResponse(200, "Successfully logged in", message)

	ctx.JSON(http.StatusOK, response)
}

//user management

func (ad *AdminHandler) GetAllUsers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add page number as params", err.Error(), "")
		ctx.JSON(400, response)

	}
	pagesize, err := strconv.Atoi(ctx.Query("pagesize"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add pages size as params", err.Error(), "")
		ctx.JSON(400, response)
	}
	pagination := utils.Pagination{
		Page:     page,
		PageSize: pagesize,
	}
	users, metadata, err := ad.AdminUsecase.FindAllUsers(ctx, pagination)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding all users", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "successfully displaying all users", users, metadata)
	ctx.JSON(http.StatusOK, response)
}

//to block user

func (ad *AdminHandler) BlockUser(ctx *gin.Context) {
	fmt.Println("handleer.....")
	var blockuser request.Block
	if err := ctx.ShouldBindJSON(&blockuser); err != nil {
		response := response.ErrorResponse(400, "error while getting id from admin", err.Error(), nil)
		ctx.JSON(400, response)
		return
	}
	err := ad.AdminUsecase.BlockUser(ctx, int(blockuser.UserID))
	if err != nil {
		response := response.ErrorResponse(400, "error while block user", err.Error(), nil)
		ctx.JSON(400, response)
		return

	}
	response := response.SuccessResponse(200, "successfully blocked user", nil, blockuser)
	ctx.JSON(http.StatusOK, response)

}
func (ad *AdminHandler) UnBlockUser(ctx *gin.Context) {
	var unblockuser request.Block
	if err := ctx.ShouldBindJSON(&unblockuser); err != nil {
		response := response.ErrorResponse(400, "error while getting id from admin", err.Error(), nil)
		ctx.JSON(400, response)
		return
	}

	err := ad.AdminUsecase.UnBlockUser(ctx, int(unblockuser.UserID))
	if err != nil {
		response := response.ErrorResponse(400, "error while unblock user", err.Error(), nil)
		ctx.JSON(400, response)
		return

	}
	response := response.SuccessResponse(200, "successfully unblocked user", nil, unblockuser)
	ctx.JSON(http.StatusOK, response)
}
