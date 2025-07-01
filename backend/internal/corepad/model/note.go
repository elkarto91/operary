package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content       string             `json:"content" bson:"content"`
	Author        string             `json:"author" bson:"author"`
	Tags          []string           `json:"tags" bson:"tags"`
	Timestamp     time.Time          `json:"timestamp" bson:"timestamp"`
	Media         []string           `json:"media,omitempty" bson:"media,omitempty"`
	PriorityScore float64            `json:"priority_score,omitempty" bson:"priority_score,omitempty"`
	SuggestAction bool               `json:"suggest_action,omitempty" bson:"suggest_action,omitempty"`
}
