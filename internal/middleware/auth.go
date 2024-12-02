package middleware

import (
	"context"
	"errors"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)

		userToken, err := getTokenEntityByToken(token)

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

func getTokenEntityByToken(token string) (entities.UserToken, error) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", err)
		}
	}()
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	userToken := entities.UserToken{}

	if err := gDB.Where("token = ? AND expiration_time > ?", token, formattedTime).First(&userToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Токен не найден или истек")
			return userToken, err
		}
	}
	return userToken, nil
}
