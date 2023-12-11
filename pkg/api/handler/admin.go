package handler

import (
	"errors"
	"glamgrove/pkg/auth"
	services "glamgrove/pkg/usecase/interfaces"

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
