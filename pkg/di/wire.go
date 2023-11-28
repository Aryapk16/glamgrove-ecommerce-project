//go:build wireinject
// +build wireinject

package di

import (
	http "glamgrove/pkg/api"
	handler "glamgrove/pkg/api/handler"
	"glamgrove/pkg/config"
	"glamgrove/pkg/db"
	userRepo "glamgrove/pkg/repository"
	userUseCase "glamgrove/pkg/usecase"

	"github.com/google/wire"
)

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatbase,
		userRepo.NewUserRepository,
		userUseCase.NewUserUseCase,
		handler.NewUserHandler,

		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
