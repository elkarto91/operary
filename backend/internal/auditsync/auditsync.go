package auditsync

import (
	"github.com/elkarto91/operary/internal/auditsync/handler"
	"github.com/elkarto91/operary/internal/auditsync/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/auditsync", func(r chi.Router) {
		r.Post("/entries", handler.CreateAuditEntry)
		r.Get("/entries", handler.GetAuditEntries)
	})
}
