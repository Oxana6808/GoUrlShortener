package handlers

import (
	"GoUrlShortener/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

// 📝 Структуры для работы с запросами и ответами
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// 📌 Обработчик создания короткого URL
func ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}

	// Генерируем короткий код
	shortCode := generateShortCode()

	// Записываем в базу данных
	url := repository.URL{ShortURL: shortCode, Original: req.URL}
	if err := repository.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения URL"})
		return
	}

	// Возвращаем клиенту сокращённый URL
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortCode)})
}

// 📌 Обработчик редиректа по короткому URL
func RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	// Ищем оригинальный URL в БД
	var url repository.URL
	if err := repository.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL не найден"})
		return
	}

	// Делаем редирект на оригинальный URL
	c.Redirect(http.StatusMovedPermanently, url.Original)
}

// 🔧 Функция генерации случайного кода для короткой ссылки
func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	short := make([]byte, 6)
	for i := range short {
		short[i] = letters[rand.Intn(len(letters))]
	}
	return string(short)
}
