FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server ./cmd/app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080

CMD ["./server"]