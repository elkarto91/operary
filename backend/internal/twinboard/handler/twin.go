package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/elkarto91/operary/internal/twinboard/usecase"
)

// TriggerSnapshot handles POST /twinboard/snapshot to capture the current state.
func TriggerSnapshot(w http.ResponseWriter, r *http.Request) {
	snap, err := usecase.GenerateSnapshot()
	if err != nil {
		http.Error(w, "Failed to generate snapshot", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(snap)
}

// GetLiveState returns the most recently generated snapshot.
func GetLiveState(w http.ResponseWriter, r *http.Request) {
	snap, err := usecase.GetLiveSnapshot()
	if err != nil {
		http.Error(w, "No snapshot available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snap)
}

// GetHistory returns snapshots between provided start and end times.
func GetHistory(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		http.Error(w, "invalid start", http.StatusBadRequest)
		return
	}
	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		http.Error(w, "invalid end", http.StatusBadRequest)
		return
	}

	snaps, err := usecase.GetHistory(start, end)
	if err != nil {
		http.Error(w, "failed to fetch history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snaps)
}
