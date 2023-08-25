.PHONY: migrate-build
migrate-build:
	go build -o build/migrate ./cmd/migrate

migrate-up: migrate-build
	./build/migrate --action=up

migrate-down: migrate-build
	./build/migrate --action=down
