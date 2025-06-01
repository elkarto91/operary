package usecase

import (
	"errors"
	"time"

	"github.com/elkarto91/operary/internal/equiptrust/model"
	"github.com/elkarto91/operary/internal/equiptrust/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EquipmentCheckoutRequest struct {
	EquipmentID string `json:"equipment_id"`
	UsedBy      string `json:"used_by"`
	Purpose     string `json:"purpose"`
	ShiftID     string `json:"shift_id"`
	Notes       string `json:"notes"`
}

func Checkout(req EquipmentCheckoutRequest) (model.EquipmentLedger, error) {
	entry := model.EquipmentLedger{
		EquipmentID: req.EquipmentID,
		UsedBy:      req.UsedBy,
		Purpose:     req.Purpose,
		ShiftID:     req.ShiftID,
		Notes:       req.Notes,
	}
	return repo.Save(entry)
}

func Return(id string, returnTime time.Time) (model.EquipmentLedger, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.EquipmentLedger{}, errors.New("invalid ID")
	}
	return repo.UpdateReturn(objID, returnTime)
}

func ListLedger() ([]model.EquipmentLedger, error) {
	return repo.List()
}
