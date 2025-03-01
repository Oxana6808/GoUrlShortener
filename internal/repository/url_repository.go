package repository

import (
	"errors"
	"fmt"
)

//// Функция сохранения короткого URL в базе
//func SaveURL(short, original string) {
//
//	// Создаём объект URL
//	url := URL{ShortURL: short, Original: original}
//
//	// Сохраняем в БД
//	result := DB.Create(&url)
//	if result.Error != nil {
//		fmt.Println("Ошибка сохранения в БД:", result.Error)
//		return // Просто завершаем выполнение, не возвращаем ошибку
//	}
//
//	fmt.Println("Успешно сохранено в БД:", url)
//}

func SaveURL(short, original string) error {
	var existing URL
	// Проверяем, существует ли запись с таким оригинальным URL
	if err := DB.Where("original = ?", original).First(&existing).Error; err == nil {
		fmt.Println("🔄 Этот URL уже есть:", existing.ShortURL)
		return nil // Если запись уже существует, возвращаем nil, чтобы не создавать дубликат
	}

	// Создаём новый объект URL
	url := URL{ShortURL: short, Original: original}
	// Сохраняем объект в базе данных
	if err := DB.Create(&url).Error; err != nil {
		fmt.Println("❌ Ошибка сохранения в БД:", err)
		return err
	}
	fmt.Println("✅ Успешно сохранено в БД:", url)
	return nil
}

// Функция поиска оригинального URL по короткому
func GetOriginalURL(short string) (string, error) {
	var url URL
	if err := DB.Where("short_url = ?", short).First(&url).Error; err != nil {
		return "", errors.New("❌ URL not found")
	}
	return url.Original, nil
}
