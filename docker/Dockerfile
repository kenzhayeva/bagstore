FROM golang:1.20

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

RUN go build -o bagstore

EXPOSE 8081

CMD ["./bagstore"]
<<<<<<< HEAD
=======


# docker/Dockerfile
#
#FROM golang:1.21 AS builder
#
#WORKDIR /app
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o bagstore main.go
#
#FROM ubuntu:22.04
#WORKDIR /app
#COPY --from=builder /app/bagstore .
#
## Устанавливаем минимально необходимые зависимости
#RUN apt-get update && \
#    apt-get install -y --no-install-recommends ca-certificates tzdata && \
#    rm -rf /var/lib/apt/lists/*
#
## Создаем непривилегированного пользователя
#RUN useradd -m appuser
#USER appuser
#
#EXPOSE 8080
#CMD ["./bagstore"]
>>>>>>> be3f14a (Initial commit)
