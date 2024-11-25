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

	r.HandleFunc("/api/targets", middleware.Auth(middleware.Logging(httphandler.TargetsList))).Methods("GET")
	r.HandleFunc("/api/targets", middleware.Auth(middleware.Logging(httphandler.CreateTargetHandler))).Methods("POST")

	r.HandleFunc("/api/targets/{id}", middleware.Auth(middleware.Logging(httphandler.TargetUpdate))).Methods("PATCH")
	r.HandleFunc("/api/targets/{id}", middleware.Auth(middleware.Logging(httphandler.TargetDelete))).Methods("DELETE")

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
