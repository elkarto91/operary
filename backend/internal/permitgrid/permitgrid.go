package permitgrid

import (
	"github.com/elkarto91/operary/internal/permitgrid/handler"
	"github.com/elkarto91/operary/internal/permitgrid/repo"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/permitgrid", func(r chi.Router) {
		r.Post("/requests", handler.CreatePermitRequest)
		r.Get("/requests", handler.ListPermitRequests)
		r.Patch("/requests/{id}/approve", handler.ApprovePermit)
		r.Patch("/requests/{id}/reject", handler.RejectPermit)
	})
}
