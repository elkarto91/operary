package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/flowgrid/model"
	"github.com/elkarto91/operary/internal/flowgrid/usecase"
)

// MatchSkills accepts a ScheduleRequest and returns best-fit assignments.
func MatchSkills(w http.ResponseWriter, r *http.Request) {
	var req model.ScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	assignments := usecase.MatchSkills(req.Workers, req.Tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignments)
}

// GenerateSchedule is an alias of MatchSkills for now.
func GenerateSchedule(w http.ResponseWriter, r *http.Request) {
	MatchSkills(w, r)
}
