package handler

import (
	"errors"
	"fmt"
	"glamgrove/pkg/auth"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"glamgrove/pkg/verify"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userUsecase interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userUsecase}
}

func (u *UserHandler) UserSignup(c *gin.Context) {
	var body request.SignupUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user domain.User
	if err := copier.Copy(&user, body); err != nil {
		fmt.Println("Copy failed")
	}
	
	usr, err := u.userService.SignUp(c, user)
	if err != nil {
		response := response.ErrorResponse(400, "User already exist", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Account created successfuly", usr)
	c.JSON(http.StatusOK, response)
}


func (u *UserHandler) LoginSubmit(c *gin.Context) {
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.Email == "" && body.Password == "" && body.UserName == "" {
		_ = errors.New("Please enter user_name and password")
		response := "Field should not be empty"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user domain.User
	copier.Copy(&user, body)
	// validate login data
	user, err := u.userService.Login(c, user)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to login", err.Error(), user)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Setup JWT
	if !auth.JwtCookieSetup(c, "user-auth", user.ID) {
		response := response.ErrorResponse(500, "Generate JWT failure", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
	}
	// Success response
	response := response.SuccessResponse(200, "OTP send to your mobile number!", user.Phone)
	c.JSON(http.StatusOK, response)
}


func (u *UserHandler) UserOTPVerify(c *gin.Context) {
	var body request.OTPVerify
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user = domain.User{
		ID: body.UserID,
	}
	usr, err := u.userService.OTPLogin(c, user)
	if err != nil {
		response := response.ErrorResponse(500, "User not registered", err.Error(), user)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Verify OTP
	err = verify.TwilioVerifyOTP("+91"+usr.Phone, body.OTP)
	if err != nil {
		response := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Setup JWT
	ok := auth.JwtCookieSetup(c, "user-auth", usr.ID)
	if !ok {
		response := response.ErrorResponse(500, "Failed to login", "", nil)
		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response := response.SuccessResponse(200, "Successfuly logged in!", nil)
	c.JSON(http.StatusOK, response)
}
