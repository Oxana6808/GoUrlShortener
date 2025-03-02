package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Настройка тестовой базы данных
func setupTestDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("❌ Ошибка при создании тестовой БД: " + err.Error())
	}
	DB.AutoMigrate(&URL{})
}

func TestSaveURL(t *testing.T) {
	setupTestDB()
	err := SaveURL("abc123", "https://google.com")
	assert.NoError(t, err)
}

func TestGetOriginalURL_NotFound(t *testing.T) {
	setupTestDB()
	_, err := GetOriginalURL("nonexistent")
	assert.Error(t, err)
}
