package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/auditsync/usecase"
	metrics "github.com/elkarto91/operary/internal/handlers"
)

func CreateAuditEntry(w http.ResponseWriter, r *http.Request) {
	var req usecase.AuditEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	audit, err := usecase.CreateAudit(req)
	if err != nil {
		http.Error(w, "Failed to save audit", http.StatusInternalServerError)
		metrics.AuditSyncSubmissionFailures.Inc()

		return
	}
	metrics.AuditSyncTotalSubmitted.Inc()

	json.NewEncoder(w).Encode(audit)
}

func GetAuditEntries(w http.ResponseWriter, r *http.Request) {
	audits, err := usecase.ListAudits()
	if err != nil {
		http.Error(w, "Failed to retrieve audits", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(audits)
}
