package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type AdminRepository interface {
	FindAdmin(c context.Context, Username string) (domain.Admin, error)
	AddAdmin(c context.Context, admin domain.Admin) (domain.Admin, error)
	FindByUsername(c context.Context, Username string) (domain.Admin, error)
}
        