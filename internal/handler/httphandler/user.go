package httphandler

import (
	"encoding/json"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/service"
	"io"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var userAuth service.UserAuth
	if err := readBody(r, &userAuth); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Аутентификация
	gDB := new(db.GormDB).Connect()
	user, err := service.Authenticate(gDB, &userAuth)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Выдача токена
	token, err := entities.FindOrCreateToken(user.ID, gDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderJSON(w, &entities.AuthResponse{
		Token: token.Token,
	})
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	gDB := new(db.GormDB).Connect()
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
	gDB := new(db.GormDB).Connect()
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
