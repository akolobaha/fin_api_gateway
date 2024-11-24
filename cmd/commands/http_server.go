package commands

import (
	"context"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/handler/httphandler"
	"fin_api_gateway/internal/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"time"
)

func RunHttp(ctx context.Context, cfg *config.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/api/ping", ping).Methods("GET")

	r.HandleFunc("/api/auth", httphandler.Auth).Methods("POST")

	r.HandleFunc("/api/users", httphandler.AddUser).Methods("POST")
	r.HandleFunc("/api/users", middleware.Auth(middleware.Logging(httphandler.UpdateUser))).Methods("PATCH")

	r.HandleFunc("/api/security-fulfils", middleware.Auth(middleware.Logging(httphandler.SecurityFulfilsList))).Methods("GET")
	r.HandleFunc("/api/security-fulfils", middleware.Auth(middleware.Logging(httphandler.CreateUserSecurityFulfilHandler))).Methods("POST")

	r.HandleFunc("/api/security-fulfils/{id}", middleware.Auth(middleware.Logging(httphandler.SecurityFulfilUpdate))).Methods("PATCH")
	r.HandleFunc("/api/security-fulfils/{id}", middleware.Auth(middleware.Logging(httphandler.SecurityFulfilDelete))).Methods("DELETE")

	go func() {
		err := http.ListenAndServe(cfg.ServerAddress, r)
		if err != nil {
			slog.Info("Error starting server:", "error", err.Error())
			// Ждем несколько секунд перед перезапуском
			time.Sleep(5 * time.Second)
			slog.Info("Error starting server:", "error", err.Error())
		}

		slog.Info("Error starting server:", "error", err.Error())
	}()

	slog.Info("Сервер http запущен")

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
