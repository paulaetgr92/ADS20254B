# Etapa 1: Build da aplicação
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ajusta conforme aonde está teu main.go
RUN go build -o server ./cmd/api

# Etapa 2: Imagem final mínima
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8081

CMD ["./server"]
