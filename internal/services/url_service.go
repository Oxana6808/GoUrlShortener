package services

import (
	"crypto/sha256"
	"encoding/base64"
)

// Генерация случайного короткого URL
func GenerateShortURL(original string) string {
	//b := make([]byte, 4) // 4 байта = 6 символов в base64
	//rand.Read(b)
	//return base64.URLEncoding.EncodeToString(b)[:6]

	hash := sha256.Sum256([]byte(original))
	return base64.URLEncoding.EncodeToString(hash[:])[:6]
}
