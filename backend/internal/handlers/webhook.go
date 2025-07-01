package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/elkarto91/operary/internal/models"
)

type WebhookEvent struct {
	MachineID string `json:"machine_id"`
	Alert     string `json:"alert"`
	ShiftID   string `json:"shift_id"`
}

// WebhookHandler accepts external alerts and creates a task
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var event WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	title := fmt.Sprintf("Alert from %s: %s", event.MachineID, event.Alert)
	task, err := models.CreateTask(title, "auto")
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	atomic.AddUint64(&tasksCreated, 1)
	_ = models.RecordAudit("task", task.ID, "create", "webhook", task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
