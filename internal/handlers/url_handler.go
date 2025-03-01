package handlers

import (
	"GoUrlShortener/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

// Структуры для работы с запросами и ответами
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// Обработчик создания короткого URL
func ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}

	// Генерируем короткий код
	shortCode := generateShortCode()

	// Сохраняем URL в БД через repository.SaveURL (без возвращения ошибки)
	repository.SaveURL(shortCode, req.URL)

	// Формируем сокращённый URL для ответа клиенту
	fullShortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
	fmt.Println("🔗 Сгенерирован короткий URL:", fullShortURL)
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: fullShortURL})
}

// Обработчик редиректа по короткому URL
func RedirectURL(c *gin.Context) {

	shortURL := c.Param("shortURL")
	fmt.Println("🔍 Проверяем shortURL:", shortURL) // Логируем код

	// Проверяем наличие в БД
	var url repository.URL
	if err := repository.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		fmt.Println("❌ Не найден в БД:", shortURL)
		c.JSON(http.StatusNotFound, gin.H{"error": "URL не найден"})
		return
	}

	fmt.Println(" Найден в БД:", url.Original)
	c.Redirect(http.StatusMovedPermanently, url.Original)

	if shortURL == "favicon.ico" || shortURL == "shorten" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid request"})
		return
	}
}

// Функция генерации случайного кода для короткой ссылки
func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	short := make([]byte, 6)
	for i := range short {
		short[i] = letters[rand.Intn(len(letters))]
	}
	return string(short)
}
