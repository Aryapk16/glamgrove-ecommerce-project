//go:build wireinject
// +build wireinject

package di

import (
	http "glamgrove/pkg/api"
	handler "glamgrove/pkg/api/handler"
	"glamgrove/pkg/config"
	"glamgrove/pkg/db"
	repo "glamgrove/pkg/repository"
	UseCase "glamgrove/pkg/usecase"

	"github.com/google/wire"
)

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatbase,
		repo.NewUserRepository, repo.NewadminRepository, repo.NewProductRepository,

		UseCase.NewUserUseCase, UseCase.NewadminUseCase, UseCase.NewProductUseCase,
		handler.NewUserHandler,handler.NewAdminHandler,handler.NewProductHandler
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
