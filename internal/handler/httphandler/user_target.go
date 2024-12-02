package httphandler

import (
	"encoding/json"
	"errors"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strconv"
)

func CreateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var newUserTarget entities.UserTarget
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUserTarget)
	err = newUserTarget.Validate()
	if err != nil {
		jsonErrorResponse(w, err, 400)
		return
	}
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", err)
		}
	}()

	err = newUserTarget.Save(r.Context(), gDB.DB)
	if err != nil {
		jsonErrorResponse(w, err, 500)
		return
	}

	if err == nil {
		renderJSON(w, newUserTarget)
	}
	return
}

func TargetsList(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int64)

	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", err)
		}
	}()
	var results []entities.UserTarget

	err := gDB.Table("user_targets").
		Select("user_targets.*").
		Where("user_id = ? AND achieved = false", userId).
		Scan(&results).Error

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	renderJSON(w, results)
}

func TargetUpdate(w http.ResponseWriter, r *http.Request) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", err)
		}
	}()
	var updUserTarget entities.UserTarget

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updUserTarget); err != nil {
		jsonErrorResponse(w, err, 400)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		jsonErrorResponse(w, err, 400)
		return
	}

	userId := r.Context().Value("userId").(int64)
	updUserTarget.UserId = userId

	var existingRecord entities.UserTarget
	if err := gDB.Where("user_id = ? AND id = ? AND achieved = false", userId, id).First(&existingRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsonErrorResponse(w, errors.New("Target not found or it has been achieved"), 404)
			return
		} else {
			jsonErrorResponse(w, err, 500)
			return
		}
	}

	result := gDB.Model(&existingRecord).Updates(
		entities.UserTarget{
			Ticker:             updUserTarget.Ticker,
			ValuationRatio:     updUserTarget.ValuationRatio,
			Value:              updUserTarget.Value,
			FinancialReport:    updUserTarget.FinancialReport,
			NotificationMethod: updUserTarget.NotificationMethod,
		})

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	renderJSON(w, existingRecord)
}

func TargetDelete(w http.ResponseWriter, r *http.Request) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", err)
		}
	}()
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		jsonErrorResponse(w, errors.New("invalid syntax"), 400)
		return
	}

	userId := r.Context().Value("userId").(int64)
	result := gDB.Where("id = ? AND user_id = ? AND achieved = false", id, userId).Delete(&entities.UserTarget{})
	if err := result.Error; err != nil {
		jsonErrorResponse(w, err, 500)
		return
	}

	renderJSON(w, []string{"ok"})
}
