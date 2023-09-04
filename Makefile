.PHONY: migrate-build
migrate-build:
	go build -o build/migrate ./cmd/migrate

migrate-up: migrate-build
	./build/migrate --action=up

migrate-down: migrate-build
	./build/migrate --action=down

.PHONY: mock-gen
mock-gen:
	mockgen -source=./internal/module/authorization/core/authorization.go -destination=./internal/module/authorization/core/mock/authorization.go -package=mock
	mockgen -source=./internal/module/customer/core/customer.go -destination=./internal/module/customer/core/mock/customer.go -package=mock
	mockgen -source=./internal/module/password/core/password.go -destination=./internal/module/password/core/mock/password.go -package=mock