package handler

import (
	"errors"
	"fmt"
	"glamgrove/pkg/auth"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
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

// func (u *UserHandler) UserSignup(c *gin.Context) {
// 	var body request.SignupUserData
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	var user domain.User
// 	if err := copier.Copy(&user, body); err != nil {
// 		fmt.Println("Copy failed")
// 	}

// 	usr, err := u.userService.SignUp(c, user)
// 	if err != nil {
// 		response := response.ErrorResponse(400, "User already exist", err.Error(), body)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	response := response.SuccessResponse(200, "Account created successfuly", usr)
// 	c.JSON(http.StatusOK, response)
// }

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
	response := response.SuccessResponse(200, "Login succecsfull....!", user.Phone)
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
		Phone: body.PhoneNumber,
	}
	usr, err := u.userService.OTPLogin(c, user)
	if err != nil {
		response := response.ErrorResponse(500, "User not registered", err.Error(), user)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(body)
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

func (u *UserHandler) UserSignup(ctxt *gin.Context) {

	var signup domain.User

	if err := ctxt.ShouldBindJSON(&signup); err != nil {

		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return

	}

	ok, err := u.userService.FindUser(ctxt, signup)

	if err != nil {
		resp := response.ErrorResponse(400, "Error while finding user", "", nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}

	if ok {
		resp := response.ErrorResponse(400, "User already exist", "", nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}

	responce, err := verify.TwilioSendOTP("+91" + signup.Phone)

	if err != nil {
		resp := response.ErrorResponse(500, "Failed to send otp", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	token, err := auth.GenerateTokenForOtp(signup)

	if err != nil {
		resp := response.ErrorResponse(500, "unable to signup", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctxt.SetCookie("_signup-cookie", token, 20*60, "", "", false, true)

	resp := response.SuccessResponse(200, responce, nil)
	ctxt.JSON(200, resp)

}

func (u *UserHandler) VerifyOtp(ctxt *gin.Context) {

	var otp request.OTPVerify

	if err := ctxt.ShouldBindJSON(&otp); err != nil {

		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return

	}

	value, err := ctxt.Cookie("_signup-cookie")
	ctxt.SetCookie("_signup-cookie", "", -1, "", "", false, true)

	if err != nil {
		resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	details, ver := auth.ValidateOtpTokens(value)

	if ver != nil {
		resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	t := verify.TwilioVerifyOTP("+91"+details.Phone, otp.OTP)

	if t != nil {
		resp := response.ErrorResponse(400, "Invalid otp", t.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}

	var signup domain.User
	copier.Copy(&signup, &details)

	fmt.Println(signup)
	fmt.Println(details)

	_, err = u.userService.SignUp(ctxt, signup)

	if err != nil {
		resp := response.ErrorResponse(400, "Invalid", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.SuccessResponse(201, "Successfully Account Created", nil)
	ctxt.JSON(201, resp)

}

func (u *UserHandler) AddAddress(c *gin.Context) {
	var body request.Address
	userId := utils.GetUserIdFromContext(c)
	body.UserID = userId
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(400, response)
		return
	}
	if err := u.userService.Addaddress(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address saved successfully", body)
	c.JSON(200, response)

}
func (u *UserHandler) UpdateAddress(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	var body request.AddressPatch
	body.UserID = userId
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(400, response)
		return
	}
	if err := u.userService.UpdateAddress(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address updated successfuly", nil)
	c.JSON(200, response)
}

func (u *UserHandler) DeleteAddress(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	addressId, err := utils.StringToUint(c.Param("adressId"))
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	if err := u.userService.DeleteAddress(c, userId, addressId); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address deleted succesfully", nil)
	c.JSON(200, response)
}

func (u *UserHandler) GetAllAddress(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	if userId == 0 {
		response := response.ErrorResponse(500, "No user detected!", "", nil)
		c.IndentedJSON(400, response)
		return
	}
	address, err := u.userService.GetAllAddress(c, userId)

	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Get all address successfully", address)
	c.IndentedJSON(200, response)
}

func (u *UserHandler) Profile(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)

	user, err := u.userService.Profile(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly got profile", user)
	c.JSON(200, response)
}

func (u *UserHandler) AddToCart(c *gin.Context) {
	var body request.AddToCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body.ProductID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		c.JSON(400, "No user id on context")
		return
	}
	if err := u.userService.SaveCartItem(c, body); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Failed to add product item in cart", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly added product item to cart ", body)
	c.JSON(200, response)

}
func (u *UserHandler) GetcartItems(c *gin.Context) {
	var page request.ReqPagination
	count, err0 := utils.StringToUint(c.Query("count"))
	if err0 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page_number, err1 := utils.StringToUint(c.Query("page_number"))
	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	page.PageNumber = page_number
	page.Count = count

	userId := utils.GetUserIdFromContext(c)
	cartItems, err := u.userService.GetCartItemsbyCartId(c, page, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Cart Items successful", cartItems)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) UpdateCart(c *gin.Context) {
	var body request.UpdateCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		response := response.ErrorResponse(400, "No user id on context", "", nil)
		c.JSON(400, response)
		return
	}
	if err := u.userService.UpdateCart(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly updated cart", body)
	c.JSON(200, response)
}
func (u *UserHandler) DeleteCartItem(c *gin.Context) {
	var body request.DeleteCartItem
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}

	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		response := response.ErrorResponse(400, "No user id context", "", nil)
		c.JSON(400, response)
		return
	}
	if err := u.userService.RemoveCartItem(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly deleted the cart item", body)
	c.JSON(200, response)

}
