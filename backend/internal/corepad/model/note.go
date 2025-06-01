package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	Tags      []string           `json:"tags" bson:"tags"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
