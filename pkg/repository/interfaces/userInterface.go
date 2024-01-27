package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type UserRepository interface {
	SaveUser(c context.Context, user domain.User) (response.UserSignUp, error)
	FindUser(c context.Context, user domain.User) (domain.User, error)
	GetUserbyID(ctx context.Context, userId uint) (domain.User, error)

	SaveAddress(ctx context.Context, userAddress request.Address) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)
	UpdateAddress(ctx context.Context, userAddress request.AddressPatch) error
	DeleteAddress(ctx context.Context, userID, addressID uint) error
	GetDefaultAddress(ctx context.Context, userId uint) (address response.Address, err error)

	SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)
	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItem) error

	//forgot password
	FindUserByPhnNum(c context.Context, phn string,) error
}
