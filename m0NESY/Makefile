# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run: build
	@./bin/msy

build:
	@go build -o bin/msy main.go

test:
	@go test -v ./...

db-down:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) down

db-up:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) up

