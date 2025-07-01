package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/configmgr/usecase"
	"github.com/gorilla/mux"
)

func SetConfigHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	var val any
	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
		return
	}
	org := r.URL.Query().Get("org_id")
	if err := usecase.SetConfig(key, org, val); err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	org := r.URL.Query().Get("org_id")
	val, err := usecase.GetConfig(key, org)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"key": key, "value": val})
}
