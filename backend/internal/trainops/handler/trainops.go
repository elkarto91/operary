package handler

import (
	"encoding/json"
	"net/http"

	metrics "github.com/elkarto91/operary/internal/handlers"
	"github.com/elkarto91/operary/internal/trainops/usecase"
)

func IngestData(w http.ResponseWriter, r *http.Request) {
	var req usecase.IngestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sample, err := usecase.Ingest(req)
	if err != nil {
		http.Error(w, "Failed to ingest", http.StatusInternalServerError)
		return
	}
	metrics.TrainOpsSamplesTotal.Inc()
	json.NewEncoder(w).Encode(sample)
}

func ListPredictions(w http.ResponseWriter, r *http.Request) {
	samples, err := usecase.List()
	if err != nil {
		http.Error(w, "Failed to list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(samples)
}

func ProvideFeedback(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID       string `json:"id"`
		Feedback string `json:"feedback"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	sample, err := usecase.AddFeedback(payload.ID, payload.Feedback)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sample)
}
