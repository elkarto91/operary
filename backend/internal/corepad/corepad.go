package corepad

import (
	"github.com/elkarto91/operary/internal/corepad/handler"
	"github.com/elkarto91/operary/internal/corepad/repo"
	"github.com/go-chi/chi/v5"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/corepad", func(r chi.Router) {
		r.Post("/notes", handler.CreateNoteHandler)
		r.Get("/notes/{id}", handler.GetNoteHandler)
	})
}

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}
