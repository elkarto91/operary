package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/elkarto91/operary/internal/eventbus/usecase"
)

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	var req usecase.PublishRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	ev, err := usecase.PublishEvent(req)
	if err != nil {
		http.Error(w, "failed to publish", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ev)
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	sev := r.URL.Query().Get("severity")
	limit, _ := strconv.Atoi(limitStr)
	events, err := usecase.ListEvents(limit, sev)
	if err != nil {
		http.Error(w, "failed to list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}
