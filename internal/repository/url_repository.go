package repository

import (
	"errors"
)

// Функция сохранения короткого URL в базе
func SaveURL(short, original string) {
	DB.Create(&URL{ShortURL: short, Original: original})
}

// Функция поиска оригинального URL по короткому
func GetOriginalURL(short string) (string, error) {
	var url URL
	if err := DB.Where("short_url = ?", short).First(&url).Error; err != nil {
		return "", errors.New("❌ URL not found")
	}
	return url.Original, nil
}
