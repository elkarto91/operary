package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/models"
	"github.com/go-chi/chi/v5"
)

// GetShiftReportHandler returns a simple text shift report
func GetShiftReportHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "shiftID")
	if id == "" {
		http.Error(w, "Missing shift ID", http.StatusBadRequest)
		return
	}

	report, err := models.GenerateShiftReport(id)
	if err != nil {
		http.Error(w, "Failed to generate report", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"report": report})
}
