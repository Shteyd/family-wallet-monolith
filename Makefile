.PHONY: migrate-build
migrate-build:
	go build -o build/migrate ./cmd/migrate

migrate-up: migrate-build
	./build/migrate --action=up

migrate-down: migrate-build
	./build/migrate --action=down

.PHONY: mock-gen
mock-gen:
	mockgen -source=./internal/domain/customer.go -destination=./internal/domain/mock/customer.go -package=mock