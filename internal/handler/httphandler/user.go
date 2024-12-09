package httphandler

import (
	"encoding/json"
	"errors"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/log"
	"fin_api_gateway/internal/service"
	"fin_api_gateway/internal/transport"
	"fmt"
	"gorm.io/gorm"
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
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
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

	log.Info(fmt.Sprintf("Пользователь зарегистрирован: %s", user.Email))

	renderJSON(w, &entities.AuthResponse{
		Token: token.Token,
	})
}

type WithCfg struct {
	cfg  *config.Config
	path string
}

func AddUserWithRabbitHandler(cfg *config.Config, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler := &WithCfg{cfg: cfg, path: path}
		handler.AddUser(w, r)
	}
}

func (h *WithCfg) AddUser(w http.ResponseWriter, r *http.Request) {
	// инициализция RabbitMQ
	rabbit := transport.New()
	rabbit.InitConn(h.cfg)
	defer rabbit.ConnClose()
	rabbit.DeclareQueue(h.cfg.RabbitQueue)

	// Инициализация БД
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()
	var newUser entities.User
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
	}

	if err := newUser.Validate(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := newUser.SetPasswordHash(); err != nil {
		log.Error("Failed to connect to RabbitMQ:", err)
		return
	}

	if err := newUser.SetEmailConfirmationToken(); err != nil {
		log.Error("Error setting email confirmation token: ", err)
		return
	}

	result := gDB.Create(&newUser)

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Отправка сообщения в RabbitMQ
	emailConfirmMsg := entities.NewEmailConfirm(newUser.Email, *newUser.EmailConfirmationToken, r.Host, h.path)
	emailConfirmMsgData, err := json.Marshal(emailConfirmMsg)
	if err != nil {
		return
	}
	rabbit.SendMsg(emailConfirmMsgData)

	renderJSON(w, entities.UserResponse{
		ID: newUser.ID,
		UserBase: entities.UserBase{
			Name:     newUser.Name,
			Email:    newUser.Email,
			Telegram: newUser.Telegram,
		},
	})
}

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("token")
	if param == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()

	var user entities.User
	user.EmailConfirmationToken = &param

	if err := gDB.Where(user).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info(fmt.Sprintf("Польозватель с email_confirmation_token = %s не найден", param))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	user.EmailConfirmationToken = nil
	user.IsActive = true

	if err := gDB.Save(&user).Error; err != nil {
		log.Info(fmt.Sprintf("Ошибка при обновлении пользователя: %s", err))
	} else {
		log.Info("Пользователь успешно обновлен: %+v\n")
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
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
