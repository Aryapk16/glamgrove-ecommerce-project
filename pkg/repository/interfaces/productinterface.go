package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	FindProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	SaveProduct(ctx context.Context, product domain.Product) error
	FindProductByID(ctx context.Context, productID uint) (product domain.Product, err error)
	UpdateProduct(ctx context.Context, product domain.Product) error
	DeleteProduct(ctx context.Context, productID uint) (domain.Product, error)

	FindBrand(ctx context.Context, brand request.Category) (request.Category, error)
	AddCategory(ctx context.Context, brand request.Category) (err error)
	GetAllBrand(ctx context.Context) (brand []response.Brand, err error)

	//product item
	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error)
}
