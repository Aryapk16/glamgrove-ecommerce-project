package handler

import (
	"time"

	"glamgrove/pkg/auth"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"glamgrove/pkg/verification"
	"net/http"

	"glamgrove/pkg/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userUseCase service.UserUseCase
}

func NewUserHandler(usecase service.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: usecase}
}

func (u *UserHandler) Login(ctx *gin.Context) {

	var loginReq request.LoginRequest

	if ctx.ShouldBindJSON(&loginReq) != nil {

		ctx.JSON(404, gin.H{
			"StatusCode": 400,
			"msg":        "Enter values Properly",
			"error":      "Cant't bind the json",
		})
		return
	}

	var user domain.User

	copier.Copy(&user, loginReq)

	user, err := u.userUseCase.Login(ctx, user)

	if err != nil {

		ctx.JSON(400, gin.H{
			"StatusCode": 400,
			"error":      err,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
	})

	//sign the token
	signedString, err := token.SignedString([]byte(config.GetJWTConfig()))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"StatusCode": 500,
			"msg":        "Error to Create JWT",
		})
	}

	ctx.SetCookie("jwt-auth", signedString, 10*60, "", "", false, true)

	ctx.JSON(200, gin.H{
		"StatusCode": 200,
		"Status":     "Login success",
	})
}

func (u *UserHandler) SignUp(ctx *gin.Context) {

	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := response.ErrorResponse(400, "error while getting admin details", err.Error(), user)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	_, err := u.userUseCase.Signup(ctx, user)
	if err != nil {
		res := response.ErrorResponse(500, "error while signup", err.Error(), user)
		ctx.JSON(500, res)
		return
	}

	if _, err := verification.SendOtp("+91" + user.Phone); err != nil {
		res := response.ErrorResponse(400, "error while sending otp", err.Error(), user)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWTPhn(user.Phone)
	if err != nil {
		response := response.ErrorResponse(400, "failed to send otp", err.Error(), "ad didn't exist")

		ctx.JSON(400, response)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Signup_Authorization", tokenString["accessToken"], 300, "/", " ", false, true)

	response := response.SuccessResponse(200, "otp send successfully", nil)
	ctx.JSON(http.StatusOK, response)

}

// var req request.SignUpReq
// var user domain.User

// if ctx.BindJSON(&req) != nil {

// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"StatusCode": 400,
// 		"msg":        "Cant't Bind The Values",
// 		"user":       req,
// 	})

// 	return
// }

// copier.Copy(&user, req)

// user, err := u.userUseCase.Signup(ctx, user)

// if err != nil {
// 	fmt.Println("Handler: Error is", err)
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"StatusCode": 400,
// 		"msg":        "Invalid Inputs",
// 		"error":      err.Error(),
// 	})
// 	return
// }

// ctx.JSON(200, gin.H{
// 	"StatusCode": 200,
// 	"msg":        "Successfully Account Created",
// 	"user":       user,
// })
// response := response.SuccessResponse(200, "Account created successfuly", gin.H{
// 	"ID":   user.ID,
// 	"user": user.FirstName + " " + user.LastName,
// })
// ctx.JSON(http.StatusOK, response)
// }

func (us *UserHandler) VerifyOTP(c *gin.Context) {

	var user domain.User

	phonenumber, err := middleware.GetPhn(c, "Signup_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), phonenumber)
		c.JSON(400, response)
		return
	}

	var body request.OtpStruct
	if err := c.ShouldBindJSON(&body); err != nil {
		res := response.ErrorResponse(400, "error while getting otp from user", err.Error(), nil)
		utils.ResponseJSON(c, res)
		return
	}

	// verifying otp

	err1 := verification.VerifyOtp("+91"+phonenumber, body.OTP)

	if err1 != nil {
		res := response.ErrorResponse(400, "error while verifying otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = us.userUseCase.VerifyOTP(phonenumber)
	if err != nil {
		response := response.ErrorResponse(400, "failed to register", err.Error(), "register failed")

		c.JSON(400, response)
		return
	}

	message := "welcome  " + user.FirstName

	if err != nil {
		response := response.ErrorResponse(400, "failed to register", err.Error(), "register failed")

		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "user registration completed  successfully", message)
	c.JSON(http.StatusOK, response)

}
