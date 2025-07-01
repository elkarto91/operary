package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type      string             `bson:"type" json:"type"`
	Source    string             `bson:"source" json:"source"`
	Content   string             `bson:"content" json:"content"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Actor     string             `bson:"actor" json:"actor"`
	Severity  string             `bson:"severity" json:"severity"`
	OrgID     string             `bson:"org_id,omitempty" json:"org_id,omitempty"`
}
