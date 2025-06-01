package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permit struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	RequestedBy string             `json:"requested_by"`
	Status      string             `json:"status"` // PENDING, APPROVED, REJECTED
	RequestedAt time.Time          `json:"requested_at"`
	ApprovedBy  string             `json:"approved_by,omitempty"`
	ApprovedAt  *time.Time         `json:"approved_at,omitempty"`
	ShiftID     string             `json:"shift_id"`
	Tags        []string           `json:"tags"`
}
