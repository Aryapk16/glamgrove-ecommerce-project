package repository

// import (
// 	"context"
// 	"errors"
// 	"glamgrove/pkg/domain"

// 	interfaces "glamgrove/pkg/repository/interfaces"

// 	"gorm.io/gorm"
// )

// type productDatabase struct {
// 	DB *gorm.DB
// }

// func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
// 	return &productDatabase{DB}
// }
// func (pd *productDatabase) FindProductById(c context.Context, productid uint) error {
// 	var product domain.Product
// 	err := pd.DB.Where("product_id=?", productid).First(&product).Error
// 	if err != nil {
// 		return errors.New("failed to find product")
// 	}
// 	return nil
// }
