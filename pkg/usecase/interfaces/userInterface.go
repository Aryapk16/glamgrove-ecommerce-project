package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/response"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.User) (usersignup response.UserSignUp, err error)
	Login(ctx context.Context, user domain.User) (domain.User, error)
	OTPLogin(ctx context.Context, user domain.User) (domain.User, error)
}
