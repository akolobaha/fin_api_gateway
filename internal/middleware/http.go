package middleware

import (
	"context"
	"fin_api_gateway/internal/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)

		userToken, err := service.GetTokenEntityByToken(token)

		if err != nil {
			w.WriteHeader(403)
			return
		}
		ctx := context.WithValue(r.Context(), "userToken", userToken.Token)
		ctx = context.WithValue(r.Context(), "userId", userToken.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))

		fmt.Println(token)
	}
}

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
