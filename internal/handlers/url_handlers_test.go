package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"GoUrlShortener/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Настройка тестовой базы данных
func setupTestDB() {
	repository.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	repository.DB.AutoMigrate(&repository.URL{})
}

// Настройка тестового роутера
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/shorten", ShortenURL)
	router.GET("/:shortURL", RedirectURL)
	return router
}

func TestShortenURL(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	body := map[string]string{"url": "https://google.com"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRedirectURL_NotFound(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
