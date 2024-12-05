package httphandler

import (
	"encoding/json"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/service"
	"io"
	"log/slog"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var userAuth service.UserAuth
	if err := readBody(r, &userAuth); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Аутентификация
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err.Error())
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err)
		}
	}()
	user, err := service.Authenticate(gDB.DB, &userAuth)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Выдача токена
	token, err := entities.FindOrCreateToken(user.ID, gDB.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Info("Пользователь зарегистрирован", "email", user.Email)

	renderJSON(w, &entities.AuthResponse{
		Token: token.Token,
	})
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err)
		}
	}()
	var newUser entities.User
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
	}

	err := newUser.Validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result := gDB.Create(&newUser)

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	renderJSON(w, entities.UserResponse{
		ID: newUser.ID,
		UserBase: entities.UserBase{
			Name:     newUser.Name,
			Email:    newUser.Email,
			Telegram: newUser.Telegram,
		},
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err.Error())
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err)
		}
	}()
	var user entities.User
	if err := readBody(r, &user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	user.ID = r.Context().Value("userId").(int64)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	result := gDB.Model(&user).Updates(user)

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	renderJSON(w, entities.UserResponse{
		ID: user.ID,
		UserBase: entities.UserBase{
			Name:     user.Name,
			Email:    user.Email,
			Telegram: user.Telegram,
		},
	})
}

func readBody(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer json.Unmarshal(body, v)

	return nil
}
