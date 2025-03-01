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
 git clone https://github.com/Oxana6808/GoUrlShortener.git
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

```sh
Invoke-WebRequest -Uri "http://localhost:8080/shorten" -Method Post -Headers @{"Content-Type"="application/json"} -Body '{"url":"https://google.com"}'
```

## 🛠 Используемые технологии
- **Go** (Gin, Gorm)
- **PostgreSQL**
- **Docker & Docker Compose**

---

## ✨ Автор
**Dmitrieva Oksana**

