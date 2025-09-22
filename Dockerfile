# Stage 1: build
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# Stage 2: final image
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .

# Opcional: se quiser logs coloridos
ENV GIN_MODE=release

CMD ["./server"]
