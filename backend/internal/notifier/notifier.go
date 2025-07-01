package notifier

import (
	"github.com/elkarto91/operary/internal/notifier/handler"
	"github.com/elkarto91/operary/internal/notifier/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init sets up database collections.
func Init(db *mongo.Database) { repo.InitMongoCollection(db) }

// RegisterRoutes exposes HTTP endpoints for the notifier service.
func RegisterRoutes(r chi.Router) {
	r.Route("/notify", func(r chi.Router) {
		r.Post("/send", handler.SendNotification)
		r.Get("/user/{id}", handler.ListUserNotifications)
	})
}
