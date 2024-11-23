package httphandler

import (
	"encoding/json"
	"errors"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateUserSecurityFulfilHandler(w http.ResponseWriter, r *http.Request) {
	var newUserSecFulfil entities.UserSecurityFulfil
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUserSecFulfil)
	err = newUserSecFulfil.Validate()
	gormDb := new(db.GormDB)
	err = newUserSecFulfil.Save(r.Context(), gormDb.Connect())
	ProcessHttp400(err, w)

	if err == nil {
		renderJSON(w, newUserSecFulfil)
	}
	return
}

func SecurityFulfilsList(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int64)

	gDB := new(db.GormDB).Connect()
	var results []entities.UserSecurityFulfil
	paginatedDB := Paginate(r)(gDB)
	err := paginatedDB.Table("user_security_fulfils").
		Select("user_security_fulfils.*").
		Where("user_id = ?", userId).
		Scan(&results).Error

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	renderJSON(w, results)
}

func SecurityFulfilUpdate(w http.ResponseWriter, r *http.Request) {
	gDB := new(db.GormDB).Connect()
	var updUserSecFulfil entities.UserSecurityFulfil

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updUserSecFulfil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		http.Error(w, "invalid syntax", 400)
		return
	}

	userId := r.Context().Value("userId").(int64)
	updUserSecFulfil.UserId = userId

	var existingRecord entities.UserSecurityFulfil
	if err := gDB.Where("user_id = ? AND id = ?", userId, id).First(&existingRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", 404)
		} else {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	result := gDB.Model(&existingRecord).Updates(
		entities.UserSecurityFulfil{
			Ticker: updUserSecFulfil.Ticker,
			PE:     updUserSecFulfil.PE,
			PBv:    updUserSecFulfil.PBv,
		})

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	renderJSON(w, existingRecord)
}

func SecurityFulfilDelete(w http.ResponseWriter, r *http.Request) {
	gDB := new(db.GormDB).Connect()
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		http.Error(w, "invalid syntax", 400)
		return
	}

	userId := r.Context().Value("userId").(int64)
	result := gDB.Where("id = ? AND user_id = ?", id, userId).Delete(&entities.UserSecurityFulfil{})
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 500)
	}

	renderJSON(w, []string{"ok"})
}
