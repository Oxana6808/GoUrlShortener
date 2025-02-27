package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Глобальная переменная для БД
var DB *gorm.DB

// Структура таблицы URL
type URL struct {
	ID       uint   `gorm:"primaryKey"`
	ShortURL string `gorm:"uniqueIndex"`
	Original string
}

// Тестовое подключение к базе (чтобы ждать, пока она будет готова)
func TestDBConnection() bool {
	dsn := "host=db user=myuser password=mypassword dbname=shortener port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("⚠️ База данных недоступна, пробуем снова...")
		return false
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	return true
}

// Подключение к базе данных
func InitDB() {
	dsn := "host=db user=myuser password=mypassword dbname=shortener port=5432 sslmode=disable"

	// Проверяем, если уже подключены, то не переподключаемся
	if DB != nil {
		fmt.Println("🔄 База данных уже инициализирована.")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к базе данных:", err)
	}
	fmt.Println("✅ Подключение к базе данных установлено!")

	// Выполняем автоматическое создание таблицы
	if err := DB.AutoMigrate(&URL{}); err != nil {
		log.Fatal("❌ Ошибка миграции базы данных:", err)
	}
	fmt.Println("✅ Таблица успешно создана или уже существует!")
}
