package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
)

type AdminRepository interface {
	FindAdmin(c context.Context, Username string) (domain.AdminDetails, error)
	AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
}
