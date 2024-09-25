package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Event - структура для события календаря
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

// toJSON - функция сериализации объекта в JSON
func toJSON(e Event, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// parseEvent - функция для парсинга данных из запроса
func parseEvent(r *http.Request) (Event, error) {
	var e Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return Event{}, err
	}

	// Простая валидация полей
	if e.Title == "" || e.StartTime.IsZero() || e.EndTime.IsZero() || e.EndTime.Before(e.StartTime) {
		return Event{}, http.ErrBodyNotAllowed
	}

	return e, nil
}

// createEventHandler - обработчик для создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEvent(r)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// В реальном приложении здесь нужно сохранять событие в базу данных
	toJSON(event, w)
}

// updateEventHandler - обработчик для обновления события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEvent(r)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Логика обновления события
	toJSON(event, w)
}

// loggingMiddleware - middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Обработчики
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)

	// Middleware логирования
	loggedMux := loggingMiddleware(mux)

	// Запуск сервера
	log.Println("Сервер запущен на порту 8080...")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}
