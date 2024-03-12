package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"glamgrove/pkg/verify"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) service.UserService {
	return &UserUseCase{userRepository: repo}
}

func (U *UserUseCase) FindUser(ctx context.Context, user domain.User) (bool, error) {
	DBUser, err := U.userRepository.FindUser(ctx, user)
	if err != nil {
		return false, err
	}

	if DBUser.ID == 0 {
		return false, nil
	}

	return true, nil

}

func (u *UserUseCase) SignUp(ctx context.Context, user domain.User) (response.UserSignUp, error) {
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return response.UserSignUp{}, err
	}
	if DBUser.ID == 0 {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return response.UserSignUp{}, err
		}
		user.Password = string(hashedPass)
		usersignup, err := u.userRepository.SaveUser(ctx, user)
		if err != nil {
			return response.UserSignUp{}, err
		}
		return usersignup, nil
	} else {
		return response.UserSignUp{}, fmt.Errorf("%v user already exists", DBUser.UserName)
	}

}

func (u *UserUseCase) Login(c context.Context, user domain.User) (domain.User, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(c, user)
	fmt.Println(DBUser.Password)
	fmt.Println(user.Password)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("User not exist")
	}
	//Check if the user blocked by admin
	if DBUser.BlockStatus {
		return user, errors.New("User blocked by admin")
	}

	// check password with hashed pass

	err = bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println(err)
		return user, errors.New("password incorrect")
	}

	return DBUser, nil

}

func (u *UserUseCase) OTPLogin(ctx context.Context, user domain.User) (domain.User, error) {
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("User not exist")
	}
	return DBUser, nil
}

func (u *UserUseCase) Addaddress(ctx context.Context, address request.Address) error {

	if err := u.userRepository.SaveAddress(ctx, address); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) UpdateAddress(ctx context.Context, address request.AddressPatch) error {
	if err := u.userRepository.UpdateAddress(ctx, address); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	if err := u.userRepository.DeleteAddress(ctx, userID, addressID); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error) {
	address, err = u.userRepository.GetAllAddress(ctx, userId)
	if err != nil {
		return address, err
	}

	if len(address) < 1 {
		return address, errors.New("no address found for this user")
	}
	return address, nil
}

func (u *UserUseCase) Profile(ctx context.Context, userId uint) (response.Profile, error) {
	user, err := u.userRepository.GetUserbyID(ctx, userId)
	if err != nil {
		return response.Profile{}, err
	}

	defaultAddress, err := u.userRepository.GetDefaultAddress(ctx, userId)
	if err != nil {
		return response.Profile{}, err
	}
	//fmt.Println(defaultAddress)

	profile := response.Profile{
		//ID:             user.ID,
		UserName:       user.UserName,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Age:            user.Age,
		Email:          user.Email,
		Phone:          user.Phone,
		DefaultAddress: defaultAddress,
	}

	return profile, nil
}

func (u *UserUseCase) SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error {
	if err := u.userRepository.SavetoCart(ctx, addToCart); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {
	cartItems, err := u.userRepository.GetCartItemsbyUserId(ctx, page, userID)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (u *UserUseCase) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	if err := u.userRepository.UpdateCart(ctx, cartUpadates); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItem) error {
	if err := u.userRepository.RemoveCartItem(ctx, DelCartItem); err != nil {
		return err
	}
	return nil
}

//forgot password

func (u *UserUseCase) SendOtpForgotPass(ctx context.Context, phn string) error {
	err := u.userRepository.FindUserByPhnNum(ctx, phn)
	if err != nil {
		return err
	}
	// Generate OTP code
	if _, err1 := verify.TwilioSendOTP("+91" + phn); err1 != nil {
		return errors.New("can't send otp")
	}
	return nil
}
