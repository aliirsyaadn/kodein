#!/usr/bin/env bash
SHELL := /bin/bash
export NOW = $(shell date +"%F %T")

export APP_PORT=3000
export APP_ALLOW_ORIGIN=http://localhost:3000
export DB_NAME=kodein
export DB_USER=kodein
export DB_PASSWORD=developmentpass
export DB_HOST=127.0.0.1
export DB_PORT=5433
export DB_SSL_MODE=disable

install:
	@echo "configuring app"
	@go mod init github.com/aliirsyaadn/kodein
	@echo -n "$(NOW) installing dependencies... "
	@go mod tidy
	@echo "done"

build:
	@echo -n "$(NOW) building app... "
	@go build -o mainapp ./cmd/mainapp/
	@echo "done"

run:
	@echo "$(NOW) starting app... "
	@go run ./cmd/mainapp/main.go

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir files/sql/schemas $$name; \
	migrate create -ext sql -dir files/sql/seeds $$name

migrate:
	@echo -n "$(NOW) migrating app... "
	@go run ./cmd/migration/main.go
	@echo "done"

migrate-down:
	@echo -n "$(NOW) down migrating app... "
	@go run ./cmd/migration/main.go -cmd=seed -down
	@go run ./cmd/migration/main.go -down
	@echo "done"

seed:
	@echo -n "$(NOW) seeding app... "
	@go run ./cmd/migration/main.go -cmd=seed
	@echo "done"

seed-down:
	@echo -n "$(NOW) down seeding app... "
	@go run ./cmd/migration/main.go -cmd=seed -down
	@echo "done"

drop:
	@echo "$(NOW) droping database... "
	migrate -source file://files/sql/schemas \
		-database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE) drop
	@echo "done"

sqlc:
	@echo -n "$(NOW) generating sqlc... "
	sqlc generate
	@echo "done"

connect_db:
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) $(DB_NAME) -h ${DB_HOST} -p $(DB_PORT)

format:
	@echo -n "$(NOW) formating code... "
	@go fmt ./...
	@echo "done"

restart: migrate-down migrate seed 

testing:
	@go test ./test/... -v

mock:
	@mockgen -source=./services/member/main.go -destination=./test/services/member/mock/member_repository.go