
package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "operary/internal/models"
)

func StartShiftHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        SupervisorID string `json:"supervisor_id"`
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

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(shift)
}

func CloseShiftHandler(w http.ResponseWriter, r *http.Request) {
    shiftID := chi.URLParam(r, "shiftID")
    if shiftID == "" {
        http.Error(w, "Missing shift ID", http.StatusBadRequest)
        return
    }

    err := models.CloseShift(shiftID)
    if err != nil {
        http.Error(w, "Failed to close shift", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Shift closed"))
}
