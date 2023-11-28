package handler

import (
	service "glamgrove/pkg/usecase/interfaces"
)

func NewHandlers(adminUseCase service.AdminUseCase, userUseCase service.UserUseCase) (*AdminHandler, *UserHandler) {

	return &AdminHandler{adminUseCase: adminUseCase},
		&UserHandler{userUseCase: userUseCase}
}
