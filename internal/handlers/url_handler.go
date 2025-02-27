package handlers

import (
	"GoUrlShortener/internal/repository"
	"GoUrlShortener/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Обработчик POST-запроса для сокращения URL
func ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"` // JSON-поле "url" обязательно
	}

	// Проверяем корректность входных данных
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ Invalid request, URL is required"})
		return
	}

	// Генерируем короткий URL
	short := services.GenerateShortURL(req.URL)

	// Сохраняем его в базе данных
	repository.SaveURL(short, req.URL)

	// Возвращаем пользователю короткий URL
	c.JSON(http.StatusOK, gin.H{"short_url": fmt.Sprintf("http://localhost:8080/%s", short)})
}

// Обработчик GET-запроса для редиректа по короткому URL
func RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL") // Получаем короткий URL из параметров пути

	// Получаем оригинальный URL из базы данных
	originalURL, err := repository.GetOriginalURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "❌ URL not found"})
		return
	}

	// Делаем редирект на оригинальный URL
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
