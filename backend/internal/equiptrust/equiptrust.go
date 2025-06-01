package equiptrust

import (
	"github.com/elkarto91/operary/internal/equiptrust/handler"
	"github.com/elkarto91/operary/internal/equiptrust/repo"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/equiptrust", func(r chi.Router) {
		r.Post("/checkout", handler.CheckoutEquipment)
		r.Patch("/return/{id}", handler.ReturnEquipment)
		r.Get("/ledger", handler.ListLedger)
	})
}
