package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/search/usecase"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	scope := r.URL.Query().Get("scope")
	res, err := usecase.Search(q, scope)
	if err != nil {
		http.Error(w, "search error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}
