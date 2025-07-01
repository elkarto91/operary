package usecase

import (
	"fmt"

	"github.com/elkarto91/operary/internal/traceboard/model"
	"github.com/elkarto91/operary/internal/traceboard/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubmitReport validates and saves an incident report.
func SubmitReport(report model.IncidentReport) (model.IncidentReport, error) {
	if len(report.Title) <= 5 {
		return model.IncidentReport{}, fmt.Errorf("title too short")
	}
	return repo.Save(report)
}

// GetReports returns all incident reports.
func GetReports() ([]model.IncidentReport, error) {
	return repo.List()
}

// GetReport returns a single report by ID string.
func GetReport(id string) (model.IncidentReport, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.IncidentReport{}, err
	}
	return repo.GetByID(objID)
}

func DeleteReport(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return repo.DeleteByID(objID)
}
