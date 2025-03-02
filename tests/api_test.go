package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"GoUrlShortener/internal/handlers"
	"GoUrlShortener/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Настройка тестовой базы данных
func setupTestDB() {
	var err error
	repository.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("❌ Ошибка при создании тестовой БД: " + err.Error())
	}
	repository.DB.AutoMigrate(&repository.URL{})
}

// Настройка тестового роутера
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/shorten", handlers.ShortenURL)
	router.GET("/:shortURL", handlers.RedirectURL)
	return router
}

func TestShortenURL(t *testing.T) {
	setupTestDB() // 🛠️ Инициализируем тестовую базу
	router := setupRouter()

	body := map[string]string{"url": "https://google.com"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRedirectURL(t *testing.T) {
	setupTestDB() // 🛠️ Инициализируем тестовую БД
	router := setupRouter()

	// Сохраняем тестовые данные в БД
	repository.SaveURL("test123", "https://example.com")

	req, _ := http.NewRequest("GET", "/test123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "https://example.com", w.Header().Get("Location"))
}

func TestRedirectURL_NotFound(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/notexist", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
