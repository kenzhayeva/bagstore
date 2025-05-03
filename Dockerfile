FROM golang:1.23.0 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main ./main.go

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/main .
COPY internal/db/migrations ./internal/db/migrations
EXPOSE 8082
CMD ["./main"]

