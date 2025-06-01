package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditEntry struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChecklistID string             `json:"checklist_id"`
	Auditor     string             `json:"auditor"`
	Timestamp   time.Time          `json:"timestamp"`
	Passed      bool               `json:"passed"`
	Findings    string             `json:"findings"`
	Tags        []string           `json:"tags"`
	ShiftID     string             `json:"shift_id"`
}
