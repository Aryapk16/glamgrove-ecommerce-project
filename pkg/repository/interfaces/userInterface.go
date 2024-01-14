package interfaces

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/response"
)

type UserRepository interface {
	SaveUser(c context.Context, user domain.User) (response.UserSignUp, error)
	FindUser(c context.Context, user domain.User) (domain.User, error)
	GetUserbyID(ctx context.Context, userId uint) (domain.User, error)
}
