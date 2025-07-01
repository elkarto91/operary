package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SensorEvent struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SourceID    string             `json:"source_id"`
	EventType   string             `json:"event_type"`
	Timestamp   time.Time          `json:"timestamp"`
	Description string             `json:"description"`
	Tags        []string           `json:"tags"`
	TaskID      *string            `json:"task_id,omitempty"`
	AuditID     *string            `json:"audit_id,omitempty"`
}
