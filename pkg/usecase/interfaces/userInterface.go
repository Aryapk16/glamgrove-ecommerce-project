package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.User) (usersignup response.UserSignUp, err error)
	FindUser(ctx context.Context, user domain.User) (bool, error)
	Login(ctx context.Context, user domain.User) (domain.User, error)
	OTPLogin(ctx context.Context, user domain.User) (domain.User, error)
	Addaddress(ctx context.Context, address request.Address) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)
	DeleteAddress(ctx context.Context, userID, addressID uint) error
	UpdateAddress(ctx context.Context, address request.AddressPatch) error

	SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error
	GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItem) error

	Profile(ctx context.Context, userId uint) (profile response.Profile, err error)
// forgot password
	SendOtpForgotPass(ctx context.Context,phn string)error
}
