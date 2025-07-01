package integrations

import (
	"github.com/elkarto91/operary/internal/integrations/handler"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init initializes resources for integrations. Currently no DB resources are used.
func Init(db *mongo.Database) {
	_ = db
}

// RegisterRoutes exposes integration related routes under /integrations.
func RegisterRoutes(r chi.Router) {
	r.Route("/integrations", func(r chi.Router) {
		r.Post("/webhook/send", handler.SendWebhook)
		r.Post("/email/send", handler.SendEmail)
	})
}
