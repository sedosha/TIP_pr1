package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/google/uuid"
)

// user представляет структуру данных для пользователя
type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// healthResponse представляет структуру ответа для /health
type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func main() {
	// Чтение порта из переменной окружения APP_PORT
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}

	mux := http.NewServeMux()

	// Обработчик для /hello
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Обработчик для /user
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user{
			ID:   uuid.NewString(), // Генерация нового UUID
			Name: "Gopher",
		})
	})

	// НОВЫЙ обработчик для /health (опциональное задание со звёздочкой)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Формируем ответ в формате JSON с текущим временем в RFC3339
		json.NewEncoder(w).Encode(healthResponse{
			Status: "ok",
			Time:   time.Now().Format(time.RFC3339), // Время в стандартном формате
		})
	})

	// Формируем адрес для прослушивания
	addr := ":" + port
	log.Printf("Server is starting on address %s...", addr)
	// Запускаем сервер
	log.Fatal(http.ListenAndServe(addr, mux))
}