package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/sensorvault/usecase"
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
