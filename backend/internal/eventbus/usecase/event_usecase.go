package usecase

import (
	"time"

	"github.com/elkarto91/operary/internal/eventbus/model"
	"github.com/elkarto91/operary/internal/eventbus/repo"
	"go.mongodb.org/mongo-driver/bson"
)

type PublishRequest struct {
	Type     string `json:"type"`
	Source   string `json:"source"`
	Content  string `json:"content"`
	Actor    string `json:"actor"`
	Severity string `json:"severity"`
	OrgID    string `json:"org_id,omitempty"`
}

func PublishEvent(req PublishRequest) (model.Event, error) {
	ev := model.Event{
		Type:      req.Type,
		Source:    req.Source,
		Content:   req.Content,
		Actor:     req.Actor,
		Severity:  req.Severity,
		OrgID:     req.OrgID,
		Timestamp: time.Now(),
	}
	return repo.InsertEvent(ev)
}

func ListEvents(limit int, severity string) ([]model.Event, error) {
	filter := bson.M{}
	if severity != "" {
		filter["severity"] = severity
	}
	if limit <= 0 {
		limit = 50
	}
	return repo.ListEvents(limit, filter)
}
