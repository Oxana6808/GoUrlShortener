package main

import (
	"GoUrlShortener/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// Ждём, пока база будет готова
	for i := 0; i < 10; i++ {
		if repository.TestDBConnection() {
			break
		}
		fmt.Println("⏳ Ожидание базы данных...")
		time.Sleep(5 * time.Second)
	}

	// Инициализация базы данных
	repository.InitDB()

	// Создаем новый роутер
	router := gin.Default()

	// Запускаем сервер на порту 8080
	fmt.Println("🚀 Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Ошибка запуска сервера:", err)
	}
}
