package usecase

import (
	"errors"

	"github.com/elkarto91/operary/internal/permitgrid/model"
	"github.com/elkarto91/operary/internal/permitgrid/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermitRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	RequestedBy string   `json:"requested_by"`
	ShiftID     string   `json:"shift_id"`
	Tags        []string `json:"tags"`
}

func CreatePermit(req PermitRequest) (model.Permit, error) {
	permit := model.Permit{
		Title:       req.Title,
		Description: req.Description,
		RequestedBy: req.RequestedBy,
		ShiftID:     req.ShiftID,
		Tags:        req.Tags,
	}
	return repo.SavePermit(permit)
}

func ListPermits() ([]model.Permit, error) {
	return repo.ListPermits()
}

func ApprovePermit(id, approver string) (model.Permit, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Permit{}, errors.New("invalid ID")
	}
	return repo.UpdatePermitStatus(objID, "APPROVED", approver)
}

func RejectPermit(id, approver string) (model.Permit, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Permit{}, errors.New("invalid ID")
	}
	return repo.UpdatePermitStatus(objID, "REJECTED", approver)
}
