package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/response"
)

type ProductUseCase interface {
	// product
	AddProduct(c context.Context, product domain.Product) (domain.Product, error)
	FindAllProducts(c context.Context, pagination utils.Pagination) ([]response.ProductResponse, utils.Metadata, error)
	SearchByCode(c context.Context, code string) (response.ProductResponse, error)
	GetProductByID(c context.Context, ProductId int) (domain.Product, error)
	//category
	AddCategory(c context.Context, category domain.Category) (domain.Category, error)
	DisplayAllCategory(c context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error)
}
