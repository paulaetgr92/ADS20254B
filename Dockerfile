# Etapa 1: build da aplicação
FROM golang:1.23 AS builder

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum primeiro (para cache das dependências)
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar a aplicação
RUN go build -o main ./main.go

# Etapa 2: imagem final
FROM debian:bullseye-slim

WORKDIR /app

# Instalar certificados SSL (caso use chamadas HTTPS - ex: Twilio, banco remoto)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copiar binário do builder
COPY --from=builder /app/main .

# Copiar variáveis de ambiente (se existir .env)
COPY .env .env

# Expor porta
EXPOSE 8080

# Comando de start
CMD ["./main"]
