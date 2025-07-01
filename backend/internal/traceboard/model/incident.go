package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncidentReport captures metadata and analysis for an operational incident.
type IncidentReport struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	RootCause       string             `json:"root_cause"`
	Recommendations string             `json:"recommendations"`
	LinkedTaskID    *string            `json:"linked_task_id,omitempty" bson:"linked_task_id,omitempty"`
	LinkedAuditID   *string            `json:"linked_audit_id,omitempty" bson:"linked_audit_id,omitempty"`
	FaultTreeJSON   string             `json:"fault_tree_json"`
	Tags            []string           `json:"tags"`
	ReportedBy      string             `json:"reported_by"`
	CreatedAt       time.Time          `json:"created_at"`
}
