
package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "operary/internal/handlers"
)

func NewRouter() http.Handler {
    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Health check
    r.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Operary API is live."))
    })

    // Task endpoints
    r.Route("/v1/tasks", func(r chi.Router) {
        r.Get("/", handlers.ListTasksHandler)
        r.Post("/", handlers.CreateTaskHandler)
    })

    // Shift endpoints
    r.Route("/v1/shifts", func(r chi.Router) {
        r.Post("/", handlers.StartShiftHandler)
        r.Post("/{shiftID}/close", handlers.CloseShiftHandler)
    })

    return r
}
