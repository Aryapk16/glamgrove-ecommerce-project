package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type UserRepository interface {
	FindUser(ctx context.Context, user domain.Users) (domain.Users, any)
	SaveUser(ctx context.Context, user domain.Users) (domain.Users, any)
	GetAllProducts(ctx context.Context) ([]domain.Product, any)
	GetProductItems(ctx context.Context, product domain.Product) ([]domain.ProductItem, any)
}
