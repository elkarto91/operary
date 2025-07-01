package usecase

import (
	"github.com/elkarto91/operary/internal/trainops/model"
	"github.com/elkarto91/operary/internal/trainops/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IngestRequest defines the payload required to create a TrainingSample.
type IngestRequest struct {
	Source  string   `json:"source"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Outcome string   `json:"outcome,omitempty"`
}

func Ingest(req IngestRequest) (model.TrainingSample, error) {
	sample := model.TrainingSample{
		Source:  req.Source,
		Content: req.Content,
		Tags:    req.Tags,
		Outcome: req.Outcome,
	}
	return repo.Save(sample)
}

func List() ([]model.TrainingSample, error) {
	return repo.List()
}

func AddFeedback(id string, fb string) (model.TrainingSample, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.TrainingSample{}, err
	}
	return repo.UpdateFeedback(objID, fb)
}
