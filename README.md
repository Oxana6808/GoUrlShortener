# GoUrlShortener - Сервис для сокращения ссылок

## 📌 Описание проекта
**GoUrlShortener** — это веб-сервис, реализованный на **Go**, который позволяет сокращать длинные ссылки и перенаправлять пользователей по коротким URL-адресам.
Использует **PostgreSQL** для хранения данных и **Gin** в качестве веб-фреймворка.

---

## 🚀 Запуск проекта

### 📋 Требования:
- **Go 1.24+**
- **Docker и Docker Compose**
- **PostgreSQL** (в Docker контейнере)

### 🛠 Установка и запуск
1️⃣ **Клонировать репозиторий**
```sh
 git clone https://github.com/ТВОЙ_GITHUB/GoUrlShortener.git
 cd GoUrlShortener
```

2️⃣ **Запустить контейнеры**
```sh
docker-compose up --build -d
```

3️⃣ **Проверить запущенные контейнеры**
```sh
docker ps
```

✅ Должны запуститься контейнеры:
- `url_shortener_db` (**PostgreSQL**)
- `url_shortener_app` (**Go-сервер**)

---

## 🔗 API

### 1️⃣ Сокращение ссылки
**POST** `/shorten`

📌 **Запрос:**
```json
{
    "url": "https://example.com"
}
```

📌 **Ответ:**
```json
{
    "short_url": "http://localhost:8080/abc123"
}
```

### 2️⃣ Перенаправление по короткому URL
**GET** `/:shortURL`

📌 **Пример запроса:**
```
GET http://localhost:8080/abc123
```

📌 **Ожидаемый результат:** редирект на `https://example.com`

---

## 📂 Структура проекта
```
url-shortener/
│── cmd/
│── internal/
│   ├── handlers/       # Обработчики API
│   ├── repository/     # Работа с БД
│── config/             # Конфигурация
│── main.go             # Точка входа
│── go.mod              # Go модули
│── Dockerfile          # Docker конфигурация
│── docker-compose.yml  # Запуск через Docker
│── README.md           # Документация
```

---

## 🛠 Используемые технологии
- **Go** (Gin, Gorm)
- **PostgreSQL**
- **Docker & Docker Compose**

---

## ✨ Автор
**Dmitrieva Oksana**

