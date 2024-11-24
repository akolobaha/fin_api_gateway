package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging - это middleware для логирования HTTP-запросов.
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Записываем время начала обработки запроса
		start := time.Now()

		// Вызываем следующий обработчик
		next(w, r)

		// Записываем время окончания обработки запроса
		duration := time.Since(start)

		// Логируем информацию о запросе
		log.Printf(
			"%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			duration,
		)
	}
}
