# Carregar variáveis do .env
include .env
export $(shell sed 's/=.*//' .env)

# Criar banco de dados localmente
createdb:
	createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

# Rodar PostgreSQL com Docker
postgres:
	@if [ ! "$$(docker ps -q -f name=ineof-db)" ]; then \
		if [ "$$(docker ps -aq -f status=exited -f name=ineof-db)" ]; then \
			docker start ineof-db; \
		else \
			docker-compose up -d postgres; \
		fi \
	fi

# Rodar SQLC para gerar código das queries
sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

# Construir a imagem do Swagger caso não exista
swag-build:
	docker build -t meu-swag -f Dockerfile.swagger .

# Rodar Swagger para gerar documentação
docs: swag-build
	docker run --rm -v $$(pwd):/code -w /code meu-swag init -g main.go -o docs/

# Executar migrações para cima
migrateup:
	migrate -path internal/db/migration -database "$(DATABASE_URL)" -verbose up

# Reverter migrações (descer)
migrationdrop:
	migrate -path internal/db/migration -database "$(DATABASE_URL)" -verbose down

# Rodar testes com cobertura
test:
	go test -v -cover ./...

# Rodar o servidor Go
server:
	go run main.go

# Remover a pasta de docs gerada pelo Swagger
clean:
	sudo rm -rf api/docs/

.PHONY: createdb postgres sqlc-gen swag-build docs migrateup migrationdrop test server clean