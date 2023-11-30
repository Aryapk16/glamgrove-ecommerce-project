package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type UserUseCase interface {
	Signup(ctx context.Context, user domain.Users) (domain.Users, any)
	Login(ctx context.Context, user domain.Users) (domain.Users, any)
	//ShowAllProducts(ctx context.Context, user domain.Users) (domain.Users, any)
}
