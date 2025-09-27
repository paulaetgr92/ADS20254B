# Etapa de build
FROM golang:1.24-alpine AS builder

# Instala dependências necessárias
RUN apk add --no-cache gcc musl-dev

# Cria diretório da aplicação
WORKDIR /app

# Copia os arquivos go mod/sum primeiro (cache das dependências)
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o projeto
COPY . .

# Compila o binário
RUN go build -o server ./main.go

# Etapa final (imagem mínima)
FROM alpine:3.20

WORKDIR /root

# Copia o binário da etapa anterior
COPY --from=builder /app/server .

# Copia o .env (se precisar dentro do container)
COPY .env .env

# Porta que a aplicação expõe
EXPOSE 8080

# Comando padrão
CMD ["./server"]
