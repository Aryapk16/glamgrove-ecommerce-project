package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type UserRepository interface {
	FindUser(ctx context.Context, user domain.User) (domain.User, any)
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, any)
	GetProductItems(ctx context.Context, product domain.Product) ([]domain.Product, any)
	UpdateSignupstatus(phone string) error
}
