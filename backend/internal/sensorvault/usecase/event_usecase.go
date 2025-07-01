package usecase

import (
	"github.com/elkarto91/operary/internal/sensorvault/model"
	"github.com/elkarto91/operary/internal/sensorvault/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SensorEventRequest struct {
	SourceID    string   `json:"source_id"`
	EventType   string   `json:"event_type"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	TaskID      *string  `json:"task_id,omitempty"`
	AuditID     *string  `json:"audit_id,omitempty"`
}

func RecordEvent(req SensorEventRequest) (model.SensorEvent, error) {
	event := model.SensorEvent{
		SourceID:    req.SourceID,
		EventType:   req.EventType,
		Description: req.Description,
		Tags:        req.Tags,
		TaskID:      req.TaskID,
		AuditID:     req.AuditID,
	}
	return repo.Save(event)
}

func GetEvents() ([]model.SensorEvent, error) {
	return repo.List()
}

func GetEvent(id string) (model.SensorEvent, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.SensorEvent{}, err
	}
	return repo.GetByID(objID)
}

func DeleteEvent(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return repo.DeleteByID(objID)
}
