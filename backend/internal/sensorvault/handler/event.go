package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/sensorvault/usecase"
	"github.com/go-chi/chi/v5"
)

func RecordEvent(w http.ResponseWriter, r *http.Request) {
	var req usecase.SensorEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	event, err := usecase.RecordEvent(req)
	if err != nil {
		http.Error(w, "Failed to record event", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(event)
}

func ListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := usecase.GetEvents()
	if err != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	event, err := usecase.GetEvent(id)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := usecase.DeleteEvent(id); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
