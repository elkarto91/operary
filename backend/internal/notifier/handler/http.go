package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/elkarto91/operary/internal/notifier/usecase"
	"github.com/go-chi/chi/v5"
)

// SendNotification handles POST /notify/send
func SendNotification(w http.ResponseWriter, r *http.Request) {
	var req usecase.NotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	notif, err := usecase.Send(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(notif)
}

// ListUserNotifications handles GET /notify/user/{id}
func ListUserNotifications(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	limitStr := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitStr)
	notifs, err := usecase.ListUserNotifications(id, limit)
	if err != nil {
		http.Error(w, "failed to list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notifs)
}
