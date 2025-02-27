package services

import (
	"crypto/rand"
	"encoding/base64"
)

// Генерация случайного короткого URL
func GenerateShortURL(original string) string {
	b := make([]byte, 4) // 4 байта = 6 символов в base64
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}
