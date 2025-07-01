package trainops

import (
	"github.com/elkarto91/operary/internal/trainops/handler"
	"github.com/elkarto91/operary/internal/trainops/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/trainops", func(r chi.Router) {
		r.Post("/ingest", handler.IngestData)
		r.Get("/predictions", handler.ListPredictions)
		r.Post("/feedback", handler.ProvideFeedback)
	})
}
