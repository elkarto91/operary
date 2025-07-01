package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notification represents a message delivered to a user via a specific channel
// such as email or push notification.
type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	Channel   string             `bson:"channel" json:"channel"`
	Message   string             `bson:"message" json:"message"`
	Status    string             `bson:"status" json:"status"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}
