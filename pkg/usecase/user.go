package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
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
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("User not exist")
	}
	//Check if the user blocked by admin
	if DBUser.BlockStatus {
		return user, errors.New("User blocked by admin")
	}

	if _, err := verify.TwilioSendOTP("+91" + DBUser.Phone); err != nil {
		return user, fmt.Errorf("Failed to send otp %v",
			err)
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password)) != nil {
		return user, errors.New("Password incorrect")
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
