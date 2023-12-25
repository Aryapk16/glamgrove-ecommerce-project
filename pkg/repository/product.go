package repository

import (
	"context"
	"errors"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/response"

	interfaces "glamgrove/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}
func (pd *productDatabase) AddProduct(c context.Context, product domain.Product) (domain.Product, error) {
	err := pd.DB.Create(&product).Error
	if err != nil {
		return domain.Product{}, errors.New("failed to add product")
	}
	return product, nil
}
func (pd *productDatabase) FindProductById(c context.Context, productid uint) error {
	var product domain.Product
	err := pd.DB.Where("product_id=?", productid).First(&product).Error
	if err != nil {
		return errors.New("failed to find product")
	}
	return nil
}
func (pd *productDatabase) FindProduct(c context.Context, product domain.Product) (domain.Product, error) {
	err := pd.DB.Where("product_id=? OR product_name=?", product.ProductID, product.ProductName).First(&product).Error
	if err != nil {
		return domain.Product{}, errors.New("failed to find product")
	}

	return product, nil
}
func (pd *productDatabase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]response.ProductResponse, utils.Metadata, error) {
	var products []response.ProductResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.Product{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []response.ProductResponse{}, utils.Metadata{}, err
	}

	query := `select product_code,product_name,product_price,product_gst from product_details limit $1 offset $2`
	err := db.Raw(query, pagination.Limit(), pagination.Offset()).Scan(&products).Error
	if err != nil {
		return []response.ProductResponse{}, utils.Metadata{}, errors.New("failed to find all products")
	}
	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return products, metadata, nil
}
func (pd *productDatabase) SearchByCode(c context.Context, code string) (response.ProductResponse, error) {
	var product response.ProductResponse
	query := `select product_name,product_price,product_gst from product_details where product_code=?`
	pd.DB.Raw(query, code).Scan(&product)
	if product.Product_Name == "" {
		return response.ProductResponse{}, errors.New("failed to find product")
	}
	return product, nil
}
func (pd *productDatabase) GetProductByID(c context.Context, productid int) (domain.Product, error) {
	var product domain.Product
	err := pd.DB.Where("product_id=?", productid).First(&product).Error
	if err != nil {
		return domain.Product{}, errors.New("failed to find product")
	}
	return product, nil
}
func (pd *productDatabase) DeleteProduct(c context.Context, productid uint) error {
	var product_details domain.Product
	err := pd.DB.Where("product_id=?", productid).Delete(&product_details).Error
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}
//category
func (pd *productDatabase) FindCategory(c context.Context, category domain.Category) (domain.Category, error) {
	var tempCategory domain.Category
	err := pd.DB.Where("category_name=?", category.CategoryName).First(&tempCategory).Error
	if err != nil {
		return domain.Category{}, errors.New("failed find category")
	}
	return tempCategory, nil
}
func (pd *productDatabase) AddCategory(c context.Context, category domain.Category) (domain.Category, error) {
	err := pd.DB.Create(&category).Error

	if err != nil {
		return domain.Category{}, errors.New("failed to add category")
	}
	return category, nil
}
func (c *productDatabase) FindAllCategory(ctx context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error) {
	var categories []domain.Category
	var totalRecords int64

	db := c.DB.Model(&domain.Category{})

	// Get the total count of records
	if err := db.Count(&totalRecords).Error; err != nil {
		return categories, utils.Metadata{}, err
	}

	// Apply pagination
	db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

	// Fetch categories
	if err := db.Find(&categories).Error; err != nil {
		return categories, utils.Metadata{}, err
	}

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return categories, metadata, nil
}
func (pd *productDatabase) FindCategoryByName(c context.Context, categoryName string) error {
	var categories domain.Category
	err := pd.DB.Where("category_name=?", categoryName).First(&categories).Error
	if err != nil {
		return errors.New("failed find category")
	}
	return nil
}
func (pd *productDatabase) GetCategoryByID(c context.Context, categoryId int) (domain.Category, error) {
	var category domain.Category
	query := `select * from categories where id=?`
	err := pd.DB.Raw(query, categoryId).Scan(&category).Error
	if err != nil {
		return domain.Category{}, errors.New("failed to find category name")
	}

	return category, nil
}

func (pd *productDatabase) DeleteCategory(c context.Context, categoryName string) error {
	var categories domain.Category
	err := pd.DB.Where("category_name=?", categoryName).Delete(&categories).Error
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}
