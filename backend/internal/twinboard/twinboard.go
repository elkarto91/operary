package twinboard

import (
	"github.com/elkarto91/operary/internal/twinboard/handler"
	"github.com/elkarto91/operary/internal/twinboard/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init sets up the MongoDB collection used for storing snapshots.
func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

// RegisterRoutes exposes HTTP endpoints for the TwinBoard module.
func RegisterRoutes(r chi.Router) {
	r.Route("/twinboard", func(r chi.Router) {
		r.Post("/snapshot", handler.TriggerSnapshot)
		r.Get("/live", handler.GetLiveState)
		r.Get("/history", handler.GetHistory)
	})
}
