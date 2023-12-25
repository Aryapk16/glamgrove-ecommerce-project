package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/response"
)

type ProductRepository interface {
	//product
	AddProduct(c context.Context, product domain.Product) (domain.Product, error)
	FindProductById(c context.Context, productid uint) error
	FindProduct(c context.Context, product domain.Product) (domain.Product, error)
	FindAllProducts(c context.Context, pagination utils.Pagination) ([]response.ProductResponse, utils.Metadata, error)
	SearchByCode(c context.Context, code string) (response.ProductResponse, error)
	GetProductByID(c context.Context, productid int) (domain.Product, error)
	DeleteProduct(c context.Context, productid uint) error
	//UpdateProduct(c context.Context, productup request.UpdateProduct) (domain.Product, error)
	//category
	FindCategory(c context.Context, category domain.Category) (domain.Category, error)
	AddCategory(c context.Context, category domain.Category) (domain.Category, error)
	FindAllCategory(c context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error)
	GetCategoryByID(c context.Context, categoryId int) (domain.Category, error)
	FindCategoryByName(c context.Context, categoryName string) error
	//category management
	DeleteCategory(c context.Context, categoryName string) error
}
