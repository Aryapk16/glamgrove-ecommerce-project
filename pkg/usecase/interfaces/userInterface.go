package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type UserUseCase interface {
	Signup(ctx context.Context, user domain.Users) (domain.Users, any)
	Login(ctx context.Context, user domain.Users) (domain.Users, any)
	ShowAllProducts(ctx context.Context) ([]domain.Product, any)                             // show all products
	GetProductItems(ctx context.Context, product domain.Product) ([]domain.ProductItem, any) // to get all product items of a specific product
}
