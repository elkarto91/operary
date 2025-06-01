package router

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/elkarto91/operary/internal/corepad"
	"github.com/elkarto91/operary/internal/handlers"
	"github.com/elkarto91/operary/internal/opsmirror"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouterWithLogger(logger *zap.SugaredLogger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	corepad.RegisterRoutes(r)
	opsmirror.RegisterRoutes(r)

	r.Use(handlers.MetricsMiddleware)
	r.Get("/v1/metrics", handlers.MetricsHandler)

	// Custom logging middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := middleware.GetReqID(r.Context())
			logger.Infow("🛰️ Incoming request",
				"method", r.Method,
				"url", r.URL.Path,
				"request_id", reqID,
				"remote_addr", r.RemoteAddr,
			)
			next.ServeHTTP(w, r)
		})
	})

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

	r.Route("/v1/audit", func(r chi.Router) {
		r.Get("/", handlers.GetAuditLogsHandler)
	})

	return r
}
