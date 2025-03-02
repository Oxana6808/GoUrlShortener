package services

import (
	"testing"
)

// Проверяем, что длина сгенерированного URL составляет 6 символов
func TestGenerateShortURL(t *testing.T) {
	url1 := GenerateShortURL("https://google.com")
	url2 := GenerateShortURL("https://example.com")

	if len(url1) != 6 {
		t.Errorf("Ожидали длину 6, но получили %d", len(url1))
	}

	if url1 == url2 {
		t.Errorf("Ожидали уникальные значения, но получили одинаковые: %s и %s", url1, url2)
	}
}
