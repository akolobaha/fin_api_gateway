package commands

import (
	"context"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/handler/httphandler"
	"fin_api_gateway/internal/log"
	"fin_api_gateway/internal/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const ConfirmEmailPath = "/api/users/confirm-email"

func RunHttp(ctx context.Context, cfg *config.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/api/ping", ping).Methods("GET")

	r.HandleFunc("/api/auth", httphandler.Auth).Methods("POST")

	r.HandleFunc("/api/users", httphandler.AddUserWithRabbitHandler(cfg, ConfirmEmailPath)).Methods("POST")
	r.HandleFunc(ConfirmEmailPath, httphandler.ConfirmEmail).Methods("GET")
	r.HandleFunc("/api/users", middleware.Auth(middleware.Logging(httphandler.UpdateUser))).Methods("PATCH")

	r.HandleFunc("/api/targets", middleware.Auth(middleware.Logging(httphandler.TargetsList))).Methods("GET")
	r.HandleFunc("/api/targets", middleware.Auth(middleware.Logging(httphandler.CreateTargetHandler))).Methods("POST")

	r.HandleFunc("/api/targets/{id}", middleware.Auth(middleware.Logging(httphandler.TargetUpdate))).Methods("PATCH")
	r.HandleFunc("/api/targets/{id}", middleware.Auth(middleware.Logging(httphandler.TargetDelete))).Methods("DELETE")

	err := http.ListenAndServe(cfg.ServerAddress, r)
	if err != nil {
		log.Error("Error starting server:", err)
		// Ждем несколько секунд перед перезапуском
		time.Sleep(5 * time.Second)
	}

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
