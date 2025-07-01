package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID string             `bson:"user_id" json:"user_id"`
	Role   string             `bson:"role" json:"role"`
	OrgID  string             `bson:"org_id,omitempty" json:"org_id,omitempty"`
}
