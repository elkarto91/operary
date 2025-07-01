package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Supplier struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	SLA  int                `json:"sla" bson:"sla"`
}

type Delivery struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SupplierID   primitive.ObjectID `json:"supplier_id" bson:"supplier_id"`
	Item         string             `json:"item" bson:"item"`
	ExpectedDate time.Time          `json:"expected_date" bson:"expected_date"`
	Delivered    bool               `json:"delivered" bson:"delivered"`
}

type Escalation struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DeliveryID primitive.ObjectID `json:"delivery_id" bson:"delivery_id"`
	Reason     string             `json:"reason" bson:"reason"`
	Timestamp  time.Time          `json:"timestamp" bson:"timestamp"`
}
