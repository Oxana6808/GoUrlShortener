package main

import (
	"GoUrlShortener/internal/handlers"
	"GoUrlShortener/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	// Ждём, пока база будет готова
	for i := 0; i < 10; i++ {
		if repository.TestDBConnection() {
			break
		}
		fmt.Println("Ожидание базы данных...")
		time.Sleep(5 * time.Second)
	}

	// Инициализация базы данных
	repository.InitDB()

	// Создаем новый роутер
	router := gin.Default()

	// Регистрируем обработчики API
	router.POST("/shorten", handlers.ShortenURL)   // Создание короткого URL
	router.GET("/:shortURL", handlers.RedirectURL) // Редирект по короткому URL

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}

	// Выводим зарегистрированные маршруты (для отладки)
	fmt.Println(" Зарегистрированные маршруты:")
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent) // Чтобы не было ошибки 404
	})

}
