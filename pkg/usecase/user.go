package usecase

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUserCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) service.UserUseCase {
	return &userUserCase{userRepo: repo}
}

func (c *userUserCase) Login(ctx context.Context, user domain.Users) (domain.Users, any) {

	dbUser, dberr := c.userRepo.FindUser(ctx, user)

	// check user found or not
	if dberr != nil {
		return user, dberr
	}

	// check user block_status user is blocked or not
	if user.BlockStatus {
		return user, map[string]string{"Error": "User Blocked By Admin"}
	}

	//check the user password with dbPassword

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return user, map[string]string{"Error": "Entered Password is wrong"}
		}
		// Handle other potential errors more specifically if needed
		return user, map[string]string{"Error": "Error comparing passwords"}
	}

	// everything is ok then return dbUser
	return dbUser, nil
}

func (c *userUserCase) Signup(ctx context.Context, user domain.Users) (domain.Users, any) {

	// validate user values
	if err := validator.New().Struct(user); err != nil {

		errorMap := map[string]string{}
		for _, er := range err.(validator.ValidationErrors) {
			errorMap[er.Field()] = "Enter This field Properly"
		}

		return user, errorMap
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return domain.Users{}, err
	}

	user.Password = string(hash)
	user, dbErr := c.userRepo.SaveUser(ctx, user)

	return user, dbErr
}

func (c *userUserCase) ShowAllProducts(ctx context.Context) ([]domain.Product, any) {

	products, err := c.userRepo.GetAllProducts(ctx)

	if err != nil {
		return nil, map[string]string{"Error": "Can't get the products"}
	}

	return products, err
}
func (c *userUserCase) GetProductItems(ctx context.Context, product domain.Product) ([]domain.ProductItem, any) {

	productsItem, err := c.userRepo.GetProductItems(ctx, product)

	if err != nil {
		return nil, map[string]string{"Error": "To get products item"}
	}

	return productsItem, nil
}
