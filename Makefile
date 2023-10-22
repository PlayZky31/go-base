.PHONY: deploy-apps build-app test generate-mock migrate-up migrate-down migrate

include .env

MIGRATION_DIR := db/migrations
ENV_PATH := .env

deploy-apps:
	@docker compose --env-file $(ENV_PATH) down
	@docker compose --env-file $(ENV_PATH) -d db
	@docker compose --env-file $(ENV_PATH) -d apps

build-app:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o ./bin/gandiwa

test:
	go fmt ./...
	go test -coverprofile coverage.cov -cover ./...
	go tool cover -func coverage.cov

generate-mock:
	mockery --all

migrate-up:
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) up

migrate-down:
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) down

migrate:
ifdef SEQ
	@migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(strip $(SEQ))
else
	@echo "SEQ is required"
endif
