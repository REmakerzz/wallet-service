FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o wallet-service cmd/main.go

FROM gcr.io/distroless/base-debian10:latest
COPY --from=builder /app/wallet-service /wallet-service

CMD ["/wallet-service"]
