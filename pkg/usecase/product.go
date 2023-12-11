package usecase

// import (
// 	"context"
// 	"errors"
// 	"glamgrove/pkg/domain"
// 	interfaces "glamgrove/pkg/repository/interfaces"
// 	ser "glamgrove/pkg/usecase/interfaces"
// 	"glamgrove/pkg/utils"
// 	"glamgrove/pkg/utils/response"
// )

// type ProductUseCase struct {
// 	productRepo interfaces.ProductRepository
// }

// func NewProductUseCase(repo interfaces.ProductRepository) ser.ProductUseCase {
// 	return &ProductUseCase{
// 		productRepo: repo,
// 	}
// }

// // product
// func (pu *ProductUseCase) AddProduct(c context.Context, product domain.Product) (domain.Product, error) {

// 	produ, err := pu.productRepo.FindProduct(c, product)
// 	product.ProductID = produ.ProductID
// 	if err == nil {

// 		return produ, errors.New("product already exist please update product")
// 	}

// 	pro, err := pu.productRepo.AddProduct(c, product)
// 	if err != nil {
// 		return domain.Product{}, err
// 	}
// 	return pro, nil
// }

// func (pu *ProductUseCase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]response.ProductResponse, utils.Metadata, error) {
// 	product, metadata, err := pu.productRepo.FindAllProducts(c, pagination)
// 	if err != nil {
// 		return []response.ProductResponse{}, utils.Metadata{}, err
// 	}
// 	return product, metadata, nil
// }

// func (pu *ProductUseCase) SearchByCode(c context.Context, code string) (response.ProductResponse, error) {
// 	product, err := pu.productRepo.SearchByCode(c, code)
// 	if err != nil {
// 		return response.ProductResponse{}, errors.New("Invalid product,product details not available")
// 	}
// 	return product, nil
// }

// func (pu *ProductUseCase) GetProductByID(c context.Context, ProductId int) (domain.Product, error) {
// 	product, err := pu.productRepo.GetProductByID(c, ProductId)

// 	if err != nil {
// 		return domain.Product{}, err
// 	}
// 	return product, nil
// }
