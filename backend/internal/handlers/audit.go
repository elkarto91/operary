
package handlers

import (
    "encoding/json"
    "net/http"

    "operary/internal/models"
)

func GetAuditLogsHandler(w http.ResponseWriter, r *http.Request) {
    entityType := r.URL.Query().Get("entity_type")
    entityID := r.URL.Query().Get("entity_id")

    logs, err := models.GetAuditLogs(entityType, entityID)
    if err != nil {
        http.Error(w, "Failed to retrieve audit logs", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(logs)
}
