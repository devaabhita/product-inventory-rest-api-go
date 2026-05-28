FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o product-inventory main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/product-inventory .
COPY .env .
EXPOSE 8080
CMD ["./product-inventory"]