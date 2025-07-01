package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/traceboard/model"
	"github.com/elkarto91/operary/internal/traceboard/usecase"
	"github.com/go-chi/chi/v5"
)

// SubmitReport handles POST /traceboard/incidents
func SubmitReport(w http.ResponseWriter, r *http.Request) {
	var report model.IncidentReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	saved, err := usecase.SubmitReport(report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(saved)
}

// ListReports handles GET /traceboard/incidents
func ListReports(w http.ResponseWriter, r *http.Request) {
	reports, err := usecase.GetReports()
	if err != nil {
		http.Error(w, "Failed to retrieve incidents", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reports)
}

// GetReport handles GET /traceboard/incidents/{id}
func GetReport(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	report, err := usecase.GetReport(id)
	if err != nil {
		http.Error(w, "Incident not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(report)
}

func DeleteReport(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := usecase.DeleteReport(id); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
