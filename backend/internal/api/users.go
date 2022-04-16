package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rreubenreyes/hearth/internal/models"
	"gorm.io/gorm"
)

type Toolkit struct {
	db *gorm.DB
}

func (t *Toolkit) create(w http.ResponseWriter, req *http.Request) {
	var user *models.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	result := t.db.Create(user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	resp, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (t *Toolkit) root(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		t.create(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func UsersServeMux(db *gorm.DB) http.Handler {
	t := &Toolkit{db: db}
	mux := http.NewServeMux()
	mux.HandleFunc("", t.root)

	return http.StripPrefix("/api/v1/users", mux)
}
