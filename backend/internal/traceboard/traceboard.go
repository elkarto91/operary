package traceboard

import (
	"github.com/elkarto91/operary/internal/traceboard/handler"
	"github.com/elkarto91/operary/internal/traceboard/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init configures the MongoDB collection for TraceBoard.
func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

// RegisterRoutes registers TraceBoard HTTP routes.
func RegisterRoutes(r chi.Router) {
	r.Route("/traceboard", func(r chi.Router) {
		r.Post("/incidents", handler.SubmitReport)
		r.Get("/incidents", handler.ListReports)
		r.Get("/incidents/{id}", handler.GetReport)
	})
}
