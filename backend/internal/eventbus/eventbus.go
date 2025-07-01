package eventbus

import (
	"github.com/elkarto91/operary/internal/eventbus/handler"
	"github.com/elkarto91/operary/internal/eventbus/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) { repo.InitMongoCollection(db) }

func RegisterRoutes(r chi.Router) {
	r.Route("/eventbus", func(r chi.Router) {
		r.Post("/publish", handler.PublishHandler)
		r.Get("/feed", handler.FeedHandler)
	})
}
