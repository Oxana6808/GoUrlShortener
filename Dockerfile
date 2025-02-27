# Используем базовый образ Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o url-shortener ./main.go

# Запускаем приложение
CMD ["/app/url-shortener"]
