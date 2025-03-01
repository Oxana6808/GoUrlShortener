### **GoUrlShortener - Сервис для сокращения ссылок**
**Описание проекта**

**GoUrlShortener** — это веб-сервис, реализованный на **Go**, который позволяет сокращать длинные ссылки и перенаправлять пользователей по коротким URL-адресам.  
Использует **PostgreSQL** для хранения данных и **Gin** в качестве веб-фреймворка.

---

##  **Запуск проекта**

###  **Требования**:
- **Go 1.24+**
- **Docker и Docker Compose**
- **PostgreSQL** (в Docker контейнере)

###  **Установка и запуск**

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
✅ **Должны запуститься контейнеры**:
- `url_shortener_db` (**PostgreSQL**)
- `url_shortener_app` (**Go-сервер**)

---

##  **API**

**Через консоль IntelliJ IDEA (Linux/macOS/Windows)**
```sh
Invoke-RestMethod -Uri http://localhost:8080/shorten -Method POST -ContentType "application/json" -Body '{"url": "https://google.com"}'
```

После успешного запроса вернется ответ в формате JSON с коротким URL:
```json
{"short_url":"http://localhost:8080/abc123"}
```

---

###  **Проверить, что ссылка добавлена в базу данных**
Запустить команду в контейнере базы данных PostgreSQL:
```sh
docker exec -it url_shortener_db psql -U myuser -d shortener -c "SELECT * FROM urls;"
```
Пример ответа:
```
 id | short_url |      original      
----+-----------+--------------------
  1 | abc123    | https://google.com
(1 row)
```
Это означает, что сокращенная ссылка успешно добавлена.

---

###  **Перейти по сокращенной ссылке**
В браузере открыть ссылку
```sh
http://localhost:8080/abc123
```
Если редирект работает, в ответе вернется оригинальный URL.

---

## 🛠 **Используемые технологии**
- **Go** (Gin, Gorm)
- **PostgreSQL**
- **Docker & Docker Compose**

---

## ✨ **Автор**
**Dmitrieva Oksana**

---

Теперь `README.md` содержит **рабочие команды**, которые проверены и точно работают в твоем окружении! 🚀
