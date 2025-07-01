package supplymesh

import (
	"github.com/elkarto91/operary/internal/supplymesh/handler"
	"github.com/elkarto91/operary/internal/supplymesh/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollections(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/supplymesh", func(r chi.Router) {
		r.Post("/suppliers", handler.AddSupplier)
		r.Get("/suppliers", handler.ListSuppliers)
		r.Post("/deliveries", handler.AddDelivery)
		r.Get("/deliveries", handler.ListDeliveries)
		r.Post("/escalate", handler.EscalateIssue)
	})
}
