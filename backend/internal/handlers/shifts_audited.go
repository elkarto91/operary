package handlers

import (
	"encoding/json"
	"net/http"
	"sync/atomic"

	"github.com/elkarto91/operary/internal/models"
	"github.com/go-chi/chi/v5"
)

func StartShiftHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SupervisorID string `json:"supervisor_id"`
		UserID       string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.SupervisorID == "" {
		http.Error(w, "Invalid supervisor ID", http.StatusBadRequest)
		return
	}

	shift, err := models.StartShift(req.SupervisorID)
	if err != nil {
		http.Error(w, "Failed to start shift", http.StatusInternalServerError)
		return
	}

	atomic.AddUint64(&shiftsStarted, 1)

	_ = models.RecordAudit("shift", shift.ID, "start", req.UserID, shift)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shift)
}

func CloseShiftHandler(w http.ResponseWriter, r *http.Request) {
	shiftID := chi.URLParam(r, "shiftID")
	if shiftID == "" {
		http.Error(w, "Missing shift ID", http.StatusBadRequest)
		return
	}

	userID := r.URL.Query().Get("user_id")

	err := models.CloseShift(shiftID)
	if err != nil {
		http.Error(w, "Failed to close shift", http.StatusInternalServerError)
		return
	}

	atomic.AddUint64(&shiftsClosed, 1)

	_ = models.RecordAudit("shift", shiftID, "close", userID, map[string]string{"status": "closed"})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shift closed"))
}
