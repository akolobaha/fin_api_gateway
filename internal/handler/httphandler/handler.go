package httphandler

import (
	"encoding/json"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func renderJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDbRequest(w http.ResponseWriter, r *http.Request, handler func(conn *db.Connection) error) {
	gDB, err := db.ConnectToDB()
	if err != nil {
		log.Error("failed to connect to db", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("failed to close db", err)
		}
	}()

	err = handler(gDB)
	if err != nil {
		log.Error("failed to handle request", err)
		jsonErrorResponse(w, err, 500)
	}
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 20
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func jsonErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{"error": err.Error()}
	json.NewEncoder(w).Encode(response)
}
