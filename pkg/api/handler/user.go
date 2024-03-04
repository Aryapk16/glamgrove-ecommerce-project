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

// @Summary Handle user login
// @Description Authenticates user login by validating input data, checking for missing or invalid entries, and setting up JWT for authentication.
// @Tags User Profile Management
// @Accept json
// @Produce json
// @Param request body   request.LoginData true "User login details"
// @Success 200 {object} response.Response 
// @Failure 400 {object} response.Response 
// @Failure 500 {object} response.Response 
// @Router /login/ [post]
func (u *UserHandler) LoginSubmit(c *gin.Context) {
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.Email == "" && body.Password == "" {
		_ = errors.New("please enter Email and password")
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

// UserSignup registers a new user.
//
// @Summary Register a new user
// @Description Registers a new user by validating input data, checking if the user already exists, sending an OTP via Twilio, generating an authentication token, and setting a signup cookie.
// @Tags User Profile Management
// @Accept json
// @Produce json
// @Param request body domain.User true "User details for registration"
// @Success 200 {object} domain.User "Successfully registered user"
// @Failure 400 {object} response.Response "Invalid input" "Error while finding user" "User already exist"
// @Failure 500 {object} response.Response "Failed to send otp" "Unable to signup"
// @Router /signup/ [post]
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

// VerifyOtp verifies OTP for user registration.
//
// @Summary   User OTP Verification
// @Description OTP Verification to user account
// @Tags User
// @Accept json
// @Produce json
// @Param input	body	request.OTPVerify	true	"inputs"
// @Success 200 {object}	response.Response{}		"Successfully logged in"
// @Failure 400 {object} response.Response{}		"Missing or Invalid entry"
// @Failure 500 {object} response.Response{}		"Failed to login"
// @Router /signup/otp/verify [post]
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
	fmt.Println(otp)

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

// AddAddress adds a new address for a user.
//
// @Summary  Add user address
// @Description Add the address of user
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT token"
// @Param input	body	request.Address	true	"inputs"
// @Success 200 {object} response.Response{}		"Address saved successfully"
// @Failure 400 {object} response.Response{}		"Missing or Invalid entry"
// @Failure 500 {object} response.Response{} 	"Something went wrong"
// @Router /profile/add-address [post]
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

// UpdateAddress updates an existing address for a user.
//
// @Summary  update address
// @Description  update the address of user
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT token"
// @Param input	body	request.AddressPatch	true	"inputs"
// @Success 200 {object}  response.Response{}		"Address updated successfully"
// @Failure 400 {object} response.Response{}		"Missing or Invalid entry"
// @Failure 500 {object} response.Response{} 	"Something went wrong"
// @Router /profile/edit-address [put]
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

// DeleteAddress deletes an address associated with a user.
//
// @Summary    Delete user addresss
// @Description   Delete the addresss of user
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT token"
// @Param input	body	request.Address	true	"inputs"
// @Success 200 {object}  response.Response		"Address deleted successfully"
// @Failure 400 {object} response.Response		"Missing or Invalid entry"
// @Failure 500 {object} response.Response	"Something went wrong"
// @Router  /profile/delete-address/:adressId    [delete]
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

// GetAllAddress retrieves all addresses associated with a user.
//
// @Summary Get all user address
// @Description Get all the addresss of user
// @Tags User
// @Accept json
// @Produce json
// @Param  input	body	request.Address	true	"inputs"
// @Success 200 {object} response.Response{}		"Get all addresses successfully"
// @Failure 400 {object} response.Response{}    "User not detected"
// @Failure 500 {object} response.Response{} 	"Something went wrong"
// @Router  /profile/get-address [get]
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

// Profile retrieves the profile information of the authenticated user.
//
// @Summary Get user profile
// @Description Retrieve user profile details from the database
// @Tags User
// @Accept json
// @Produce json
// @Param   Authorization header string true "token"
// @Success 200 {object} response.Response{} "Successfuly got profile"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /profile/ [get]
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

// AddToCart adds a product item to the user's cart.
//
// @Summary Add a product  to cart
// @Description Add a product item to the user's cart
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token token"
// @Param body body request.AddToCartReq true "Product details to be added to cart"
// @Success 200 {object} response.Response{} "Successfuly added product item to cart"
// @Failure 400 {object} response.Response{} "Invalid input or failed to add product item to cart"
// @Router  /cart/add [post]

func (u *UserHandler) AddToCart(c *gin.Context) {
	var body request.AddToCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body.ProductItemID)
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

// GetcartItems retrieves the cart items associated with the user.
//
// @Summary Get user's cart items
// @Description Retrieve cart items of the user from the database
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true " token"
// @Param count query int true "Number of items to retrieve"
// @Param page_number query int true "Page number for pagination"
// @Success 200 {object} response.Response{} "Get Cart Items successful"
// @Failure 400 {object} response.Response{} "Missing or invalid inputs"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /cart/get [get]
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

// UpdateCart updates the user's cart.
//
// @Summary Update user's cart
// @Description  Update cart items of the user in the database
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param input body request.UpdateCartReq true "Cart update details"
// @Success 200 {object} response.Response{} "Successfuly updated cart"
// @Failure 400 {object} response.Response{} "invalid input"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /cart/update [put]
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

// DeleteCartItem deletes a cart item for the user.
//
// @Summary Delete user's cart
// @Description  Delete cart items of the user in the database
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true " token"
// @Param input body request.DeleteCartItemReq true "Cart delete details"
// @Success 200 {object} response.Response{} "Successfuly removed item from cart"
// @Failure 400 {object} response.Response{} "invalid input"
// @Failure 500 {object}  response.Response{} "Something went wrong!"
// @Router  /cart/delete [delete]

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

// forgot password
func (u *UserHandler) SendOtpForgotPass(c *gin.Context) {
	var phn request.Phn

	if err := c.ShouldBindJSON(&phn); err != nil {
		res := response.ErrorResponse(400, "error wwhile getting data from the user side", err.Error(), phn)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	phnNo := phn.Phone
	err := u.userService.SendOtpForgotPass(c, phnNo)
	if err != nil {
		res := response.ErrorResponse(400, "error while sending otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWTPhn(phnNo)
	if err != nil {
		response := response.ErrorResponse(400, "failed to send otp", err.Error(), "user didn't exist")
		c.JSON(400, response)
		return
	}
	//set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Phone_Authorization", tokenString["accessToken"], 3600*24*30, "/", " ", false, true)

	response := response.SuccessResponse(200, "otp send successfully", nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) VerifyOTPForgotPass(c *gin.Context) {
	var otp request.OTPVerify

	if err := c.ShouldBindJSON(&otp); err != nil {

		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return

	}

	value, err := c.Cookie("_forgot-cookie")
	// c.SetCookie("_forgot-cookie", "", -1, "", "", false, true)

	// if err != nil {
	// 	resp := response.ErrorResponse(500, "unable to get details", err.Error(), nil)
	// 	c.JSON(http.StatusInternalServerError, resp)
	// 	return
	// }

	details, _ := auth.ValidateOtpTokens(value)

	// if ver != nil {
	// 	resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
	// 	c.JSON(http.StatusInternalServerError, resp)
	// 	return
	// }

	t := verify.TwilioVerifyOTP("+91"+details.Phone, otp.OTP)

	if t != nil {
		resp := response.ErrorResponse(400, "Invalid otp", t.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var signup domain.User
	copier.Copy(&signup, &details)

	fmt.Println(signup)
	fmt.Println(details)

	_, err = u.userService.SignUp(c, signup)

	if err != nil {
		resp := response.ErrorResponse(400, "Invalid", err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.SuccessResponse(201, "Successfully Account Created", nil)
	c.JSON(201, resp)

}
