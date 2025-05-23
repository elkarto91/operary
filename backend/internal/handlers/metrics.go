
package handlers

import (
    "fmt"
    "net/http"
    "sync/atomic"
    "time"
)

var (
    startTime       = time.Now()
    totalRequests   uint64
)

func MetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        atomic.AddUint64(&totalRequests, 1)
        next.ServeHTTP(w, r)
    })
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
    uptime := time.Since(startTime).Seconds()
    count := atomic.LoadUint64(&totalRequests)

    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "operary_uptime_seconds %.2f
", uptime)
    fmt.Fprintf(w, "operary_total_requests %d
", count)
}
