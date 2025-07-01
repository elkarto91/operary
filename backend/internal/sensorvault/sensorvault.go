package sensorvault

import (
	"github.com/elkarto91/operary/internal/sensorvault/handler"
	"github.com/elkarto91/operary/internal/sensorvault/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/sensorvault", func(r chi.Router) {
		r.Post("/events", handler.RecordEvent)
		r.Get("/events", handler.ListEvents)
		r.Get("/events/{id}", handler.GetEvent)
		r.Delete("/events/{id}", handler.DeleteEvent)
	})
}
