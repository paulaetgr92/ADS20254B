# Etapa 1: Build da aplicação
FROM golang:1.24 AS builder

WORKDIR /app

# Copia arquivos de dependências e faz download
COPY go.mod go.sum ./
RUN go mod tidy

# Copia todo o código da aplicação
COPY . .

# Build do projeto (ajuste ./ se seu main.go não estiver na raiz)
RUN go build -o server ./

# Etapa 2: Imagem final mínima
FROM debian:bookworm-slim

WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/server .

# Porta que a aplicação vai expor
EXPOSE 8081

# Comando para rodar a aplicação
CMD ["./server"]
