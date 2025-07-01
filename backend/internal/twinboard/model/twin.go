package model

import "time"

// TwinSnapshot represents an operational state at a given moment.
type TwinSnapshot struct {
	Timestamp     time.Time         `json:"timestamp" bson:"timestamp"`
	ActiveWorkers []string          `json:"active_workers" bson:"active_workers"`
	ActiveTasks   []string          `json:"active_tasks" bson:"active_tasks"`
	EquipmentUsed map[string]string `json:"equipment_used" bson:"equipment_used"`
	Warnings      []string          `json:"warnings" bson:"warnings"`
	StatusScore   float64           `json:"status_score" bson:"status_score"`
}
