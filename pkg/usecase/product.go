package usecase

import (
	"context"
	"errors"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/response"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) service.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
	}
}

// product
func (pu *productUseCase) AddProduct(c context.Context, product domain.Product) (domain.Product, error) {

	produ, err := pu.productRepo.FindProduct(c, product)
	product.ProductID = produ.ProductID
	if err == nil {

		return produ, errors.New("product already exist please update product")
	}

	pro, err := pu.productRepo.AddProduct(c, product)
	if err != nil {
		return domain.Product{}, err
	}
	return pro, nil
}

func (pu *productUseCase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]response.ProductResponse, utils.Metadata, error) {
	product, metadata, err := pu.productRepo.FindAllProducts(c, pagination)
	if err != nil {
		return []response.ProductResponse{}, utils.Metadata{}, err
	}
	return product, metadata, nil
}

func (pu *productUseCase) SearchByCode(c context.Context, code string) (response.ProductResponse, error) {
	product, err := pu.productRepo.SearchByCode(c, code)
	if err != nil {
		return response.ProductResponse{}, errors.New("Invalid product,product details not available")
	}
	return product, nil
}

func (pu *productUseCase) GetProductByID(c context.Context, ProductId int) (domain.Product, error) {
	product, err := pu.productRepo.GetProductByID(c, ProductId)

	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

//category

func (pu *productUseCase) AddCategory(c context.Context, category domain.Category) (domain.Category, error) {

	_, err := pu.productRepo.FindCategory(c, category)

	if err == nil {
		return domain.Category{}, errors.New("category already exists")
	}
	pu.productRepo.AddCategory(c, category)

	return category, nil
}

func (pu *productUseCase) DisplayAllCategory(c context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error) {

	categories, metadata, err := pu.productRepo.FindAllCategory(c, pagination)
	if err != nil {
		return []domain.Category{}, utils.Metadata{}, errors.New("error while finding all categories")
	}
	return categories, metadata, nil
}
