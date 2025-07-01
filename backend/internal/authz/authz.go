package authz

import (
	"net/http"
	"strings"

	"github.com/elkarto91/operary/internal/authz/handler"
	"github.com/elkarto91/operary/internal/authz/repo"
	"github.com/elkarto91/operary/internal/authz/usecase"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var routeRoles = make(map[string][]string)

func Init(db *mongo.Database) {
	repo.InitMongoCollection(db)
}

func RegisterRoutes(r chi.Router) {
	r.Route("/authz", func(r chi.Router) {
		r.Post("/roles", handler.AssignRoleHandler)
		r.Get("/roles/{user}", handler.GetRoleHandler)
	})
}

func EnforceRouteRoles(route string, roles []string) {
	routeRoles[route] = roles
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Only check exact paths registered
		if allowed, ok := routeRoles[path]; ok {
			userID := r.Header.Get("X-User-ID")
			if userID == "" {
				http.Error(w, "missing user", http.StatusUnauthorized)
				return
			}
			role, _, err := usecase.GetRole(userID)
			if err != nil {
				http.Error(w, "role not found", http.StatusForbidden)
				return
			}
			for _, ar := range allowed {
				if strings.EqualFold(ar, role) {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
