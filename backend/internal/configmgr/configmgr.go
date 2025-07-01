package configmgr

import (
	"github.com/elkarto91/operary/internal/configmgr/handler"
	"github.com/elkarto91/operary/internal/configmgr/repo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Database) { repo.InitMongoCollection(db) }

func RegisterRoutes(r chi.Router) {
	r.Route("/config", func(r chi.Router) {
		r.Get("/{key}", handler.GetConfigHandler)
		r.Post("/{key}", handler.SetConfigHandler)
	})
}
