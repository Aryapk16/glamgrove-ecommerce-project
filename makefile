.PHONY: wire run swag build test mock

wire:
	cd pkg/di && wire

run:
	go run cmd/api/main.go

swag:
	swag init -g ./cmd/api/main.go ./

build:
	go build cmd/api/main.go

test: ## Run testing
	go test ./...

mock:
	mockgen -source=pkg/usecase/interfaces/userInterface.go -destination=pkg/usecase/mock/user_mock.go -package=mock
	mockgen -source=pkg/usecase/interfaces/adminInterface.go -destination=pkg/usecase/mock/admin_mock.go -package=mock
	mockgen -source=pkg/repository/interfaces/adminInterface.go -destination=pkg/repository/mock/admin_mock.go -package=mock
	mockgen -source=pkg/repository/interfaces/userInterface.go -destination=pkg/repository/mock/user_mock.go -package=mock
