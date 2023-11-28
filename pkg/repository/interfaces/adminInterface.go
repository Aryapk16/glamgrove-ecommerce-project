package interfaces

import "context"


type AdminRepository interface {
	FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, any)
	FindAllUser(ctx context.Context) ([]domain.Users, error)
	AddCategory(ctx context.Context, productCategory domain.Category) (domain.Category, any)
}
