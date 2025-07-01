package search

import (
	"github.com/elkarto91/operary/internal/search/handler"
	"github.com/go-chi/chi/v5"
)

func Init(db interface{}) {}

func RegisterRoutes(r chi.Router) {
	r.Get("/search", handler.SearchHandler)
}
