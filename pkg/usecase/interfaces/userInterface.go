package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type UserUseCase interface {
	Signup(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, user domain.User) (domain.User, any)
	VerifyOTP(phone string) error
	//ShowAllProducts(ctx context.Context, user domain.Users) (domain.Users, any)
}
