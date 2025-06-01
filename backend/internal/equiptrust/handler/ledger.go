package handler

import (
	"encoding/json"
	"net/http"

	metrics "github.com/elkarto91/operary/internal/handlers"
	"github.com/elkarto91/operary/internal/opsmirror/usecase"
)

func GetStatusDashboard(w http.ResponseWriter, r *http.Request) {
	statuses, err := usecase.FetchStatuses()
	if err != nil {
		http.Error(w, "Failed to retrieve system statuses", http.StatusInternalServerError)
		return
	}

	warnCount := 0
	for _, s := range statuses {
		if s.State == "WARN" {
			warnCount++
		}
	}

	metrics.OpsMirrorWarnings.Set(float64(warnCount))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}
