# Этап сборки: компиляция приложения
FROM golang:1.23 AS builder
WORKDIR /app

# Копируем файлы модулей и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник для Linux x86_64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o sms .

# Финальный образ на базе Alpine для минимального размера
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
# Копируем собранный бинарник и конфигурационный файл (если требуется)
COPY --from=builder /app/mail .
COPY --from=builder /app/config.yaml ./config.yaml

# Точка входа
ENTRYPOINT ["./mail"]
