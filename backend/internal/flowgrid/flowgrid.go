package flowgrid

import (
	"github.com/elkarto91/operary/internal/flowgrid/handler"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init initializes resources for FlowGrid. No DB setup yet.
func Init(db *mongo.Database) {
	_ = db // placeholder for future persistence
}

// RegisterRoutes exposes FlowGrid HTTP routes.
func RegisterRoutes(r chi.Router) {
	r.Route("/flowgrid", func(r chi.Router) {
		r.Post("/schedule", handler.GenerateSchedule)
		r.Post("/match-skills", handler.MatchSkills)
	})
}
