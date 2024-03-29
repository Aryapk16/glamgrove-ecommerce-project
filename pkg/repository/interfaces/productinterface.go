package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"time"
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

	AddImage(c context.Context, pid int, filename string) (domain.ProductImage, error)
	AddItemImage(c context.Context, pid int, filename string) (domain.ProductItemImage, error)

	//product item
	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error)

	//sales
	SalesData(sDate, Edate time.Time) (response.SalesResponse, error)

	//verify and clearcart
	DeleteCart(c context.Context, usr_id uint) error
	UpdateStatusRazorpay(c context.Context, order_id uint, order_status string, payment_status string) (response.OrderResponse, error)
}
