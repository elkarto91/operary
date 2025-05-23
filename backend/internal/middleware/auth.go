
package middleware

import (
    "net/http"
    "os"
)

// TokenAuthMiddleware checks for X-Org-Token in headers and validates it
func TokenAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        expectedToken := os.Getenv("ORG_TOKEN")
        providedToken := r.Header.Get("X-Org-Token")

        if expectedToken == "" {
            http.Error(w, "Server misconfigured: ORG_TOKEN not set", http.StatusInternalServerError)
            return
        }

        if providedToken == "" || providedToken != expectedToken {
            http.Error(w, "Unauthorized – invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
