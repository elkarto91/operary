package opsmirror

import (
	"github.com/elkarto91/operary/internal/opsmirror/handler"
	"github.com/elkarto91/operary/internal/opsmirror/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/opsmirror", func(r chi.Router) {
		r.Get("/dashboard", handler.GetStatusDashboard)
	})
}
