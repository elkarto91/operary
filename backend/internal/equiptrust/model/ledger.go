package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EquipmentLedger struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EquipmentID  string             `json:"equipment_id"`
	UsedBy       string             `json:"used_by"`
	Purpose      string             `json:"purpose"`
	ShiftID      string             `json:"shift_id"`
	CheckedOutAt time.Time          `json:"checked_out_at"`
	ReturnedAt   *time.Time         `json:"returned_at,omitempty"`
	Notes        string             `json:"notes"`
}
