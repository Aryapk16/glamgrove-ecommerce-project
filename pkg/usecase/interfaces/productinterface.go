package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"mime/multipart"
	"time"
)

type ProductService interface {
	// product

	AddProduct(ctx context.Context, product domain.Product) error
	AddCategory(ctx context.Context, Category request.Category) error
	GetAllBrands(ctx context.Context) (brand []response.Brand, err error)
	GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	UpdateProduct(ctx context.Context, product domain.Product) error
	DeleteProduct(ctx context.Context, productID uint) (domain.Product, error)
	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, count int, err error)
	AddImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.ProductImage, error)
	AddItemImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.ProductItemImage, error)

	SalesData(sDate, Edate time.Time) (response.SalesResponse, error)
	//verify and clear cart
	DeleteCart(c context.Context, usr_id uint) error
	UpdateStatusRazorpay(c context.Context, order_id uint) (response.OrderResponse, error)
}
