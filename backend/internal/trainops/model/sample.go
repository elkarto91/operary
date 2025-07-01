package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TrainingSample represents one labeled piece of historical content.
type TrainingSample struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Source    string             `json:"source"`
	Content   string             `json:"content"`
	Tags      []string           `json:"tags"`
	Outcome   string             `json:"outcome,omitempty"`
	Feedback  *string            `json:"feedback,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
}
