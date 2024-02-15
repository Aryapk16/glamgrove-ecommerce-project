//go:build wireinject
// +build wireinject

package di

import (
	http "glamgrove/pkg/api"
	"glamgrove/pkg/api/handler"
	"glamgrove/pkg/config"
	"glamgrove/pkg/db"
	"glamgrove/pkg/repository"
	"glamgrove/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(

		//Repositories
		repository.NewAdminRepository,
		repository.NewUserRepository,
		repository.NewProductRepository,
		repository.NewPaymentRepository,
		repository.NewOrderRepository,
		repository.NewCouponRepository,

		db.ConnectDatabase,

		//Usecase
		usecase.NewAdminService,
		usecase.NewUserUseCase,
		usecase.NewProductUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewOrderUseCase,
		usecase.NewCouponUseCase,

		//Handler
		handler.NewAdminHandler,
		handler.NewUserHandler,
		handler.NewProductHandler,
		handler.NewPaymentHandler,
		handler.NewOrderHandler,
		handler.NewCouponHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
