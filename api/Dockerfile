# Etapa 1: Construção da aplicação
FROM golang:1.22 AS builder

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências e baixar os pacotes
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte da aplicação
COPY . .

# Construir o binário **estático**
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/ineof-api main.go

# Etapa 2: Criar container final otimizado
FROM alpine:latest

# Instalar dependências mínimas
RUN apk --no-cache add ca-certificates bash

# Definir diretório de trabalho
WORKDIR /app

# Copiar o binário compilado da primeira etapa
COPY --from=builder /app/ineof-api /app/ineof-api

# Permitir execução do binário
RUN chmod +x /app/ineof-api

# Definir variável de ambiente para evitar buffer no output de logs
ENV GIN_MODE=release

# Definir variável de ambiente para a porta da API
ENV API_PORT=8080

# Expor a porta da API
EXPOSE 8080

# Copiar o arquivo .env para o contêiner
COPY .env /app/.env

# Comando para rodar o servidor
CMD ["/app/ineof-api"]
