FROM golang:1.20-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /room

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /room/main .
COPY rooms.db .

CMD ["./main"]
