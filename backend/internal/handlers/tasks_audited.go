package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title      string `json:"title"`
		AssignedTo string `json:"assigned_to"`
		UserID     string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task, err := models.CreateTask(req.Title, req.AssignedTo)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	_ = models.RecordAudit("task", task.ID, "create", req.UserID, task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler handles partial updates to a task's status or notes. It is
// mapped to PATCH /v1/tasks/{taskID}.
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	if taskID == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Status string `json:"status"`
		Notes  string `json:"notes"`
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task, err := models.UpdateTask(taskID, req.Status, req.Notes)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	_ = models.RecordAudit("task", taskID, "update", req.UserID, task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
