package repository

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

func (c *userDatabase) FindUser(ctx context.Context, user domain.User) (domain.User, any) {

	// Check user already exist or not
	c.DB.Raw("SELECT * FROM users where id=? OR email=? OR phone=?", user.ID, user.Email, user.Phone).Scan(&user)

	// if given userid then check mail is stil there otherwise phone or id
	if user.ID == 0 || user.Email == "" || user.Phone == "" {
		return domain.User{}, map[string]string{"Error": "Can't find the user"}
	}
	// if found the user then return user with nil
	return user, nil
}

func (c *userDatabase) SaveUser(ctx context.Context, user domain.User) (domain.User, error) {

	// check whether user is already exist
	res := c.DB.Raw("SELECT * FROM User WHERE email=? OR phone=?", user.Email, user.Phone)

	//if exist then return message as user exist
	if res.RowsAffected != 0 {
		return user, errors.New("User Already Exist")
	}

	// return user with save status
	return user, c.DB.Save(&user).Error
}

func (c *userDatabase) GetAllProducts(ctx context.Context) ([]domain.Product, any) {

	var products []domain.Product

	return products, c.DB.Raw("SELECT * FROM products").Scan(&products).Error
}
func (c *userDatabase) GetProductItems(ctx context.Context, product domain.Product) ([]domain.Product, any) {

	var productItems []domain.Product

	return productItems, c.DB.Raw("SELECT * FROM product_items WHERE product_id", product.ProductID).Scan(&productItems).Error
}

func (c *userDatabase) UpdateSignupstatus(phone string) error {

	fmt.Println(phone)

	if err := c.DB.Exec(`update users set status='Verified' where  phone=$1`, phone).Error; err != nil {
		println("eroooor", err)
		return errors.New("error while updating status " + err.Error())
	}

	return nil

}
