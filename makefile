
run: 
	air
.PHONY: run

swagger-init:
	swag init -g ./cmd/server/main.go
.PHONY: swagger-init

up: 
	docker compose --profile server up 
.PHONY: up

down:
	docker compose --profile server down 
	# docker rmi client-server:latest
.PHONY: down

lint:
	golangci-lint run --fix
	golangci-lint fmt
.PHONY: lint

test:
	go test -v ./...
.PHONY: test

migrations:
	migrate create -ext sql -dir migrations/ -seq $(FILE_NAME)
.PHONY: migrations

migrate:
	docker compose up -d
	@echo "Waiting for database to be healthy..."
	@docker compose exec postgres sh -c 'until pg_isready -U admin; do sleep 1; done'
	@echo "Database is ready!"
	migrate -path ./migrations -database "postgres://admin:admin@localhost:5432/client?sslmode=disable" up
.PHONY: migrate

migrate-down:
	docker compose up -d
	@echo "Waiting for database to be healthy..."
	@docker compose exec postgres sh -c 'until pg_isready -U admin; do sleep 1; done'
	@echo "Database is ready!"
	migrate -path ./migrations -database "postgres://admin:admin@localhost:5432/client?sslmode=disable" down -all
.PHONY: migrate-down

wire:
	wire ./cmd/server
.PHONY: wire
