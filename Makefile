.PHONY: help db-up db-down db-logs db-shell db-create run test tidy build clean unit-test-api

APP_NAME=api_biblioteca
DB_CONTAINER=postgres
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=postgres
POSTGRES_VERSION=16.14

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

#****************************************************#
# Ajuda

help:
	@echo "Comandos disponíveis:"
	@echo "  make db-up         - Sobe o container PostgreSQL"
	@echo "  make db-down       - Remove o container PostgreSQL"
	@echo "  make db-logs       - Mostra os logs do PostgreSQL"
	@echo "  make db-shell      - Conecta no psql do container"
	@echo "  make db-create     - Cria o banco da aplicação"
	@echo "  make run           - Executa a API"
	@echo "  make test          - Executa os testes"
	@echo "  make tidy          - Organiza dependências Go"
	@echo "  make build         - Compila a aplicação"
	@echo "  make clean         - Remove arquivos gerados"
	@echo "  make migrate-all   - Executa migrate completo"
	@echo "  make migrate-down  - Remove a última migration aplicada"
	@echo "  make unit-test-api - Executa os testes unitários de ./internal/api/"

#****************************************************#
# Database

db-up:
	docker run -d \
		--name $(DB_CONTAINER) \
		-p $(DB_PORT):5432 \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		postgres:$(POSTGRES_VERSION)

db-down:
	docker rm -f $(DB_CONTAINER)

db-logs:
	docker logs -f $(DB_CONTAINER)

db-shell:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER)

db-create:
	docker exec -it $(DB_CONTAINER) createdb -U $(DB_USER) $(DB_NAME)

#****************************************************#
# Migrate

migrate-all:
	migrate -path ./migration -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migration -database "$(DATABASE_URL)" down 1

migrate-seed:
	psql "$(DATABASE_URL)" -f ./migration/seed/init.sql

#****************************************************#
# Build

run:
	go run ./cmd/api

build:
	go build -o dist/$(APP_NAME) ./cmd/api

clean:
	rm -rf dist/*

#****************************************************#
# Tests

unit-test-api:
	chmod +x scripts/unit-test-api.sh
	./scripts/unit-test-api.sh

#****************************************************#
