package handler

import (
	"encoding/json"
	"net/http"

	metrics "github.com/elkarto91/operary/internal/handlers"
	"github.com/elkarto91/operary/internal/permitgrid/usecase"
	"github.com/go-chi/chi/v5"
)

func CreatePermitRequest(w http.ResponseWriter, r *http.Request) {
	var req usecase.PermitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	permit, err := usecase.CreatePermit(req)
	if err != nil {
		http.Error(w, "Failed to create permit", http.StatusInternalServerError)
		return
	}
	metrics.PermitGridTotalSubmitted.Inc()

	json.NewEncoder(w).Encode(permit)
}

func ListPermitRequests(w http.ResponseWriter, r *http.Request) {
	permits, err := usecase.ListPermits()
	if err != nil {
		http.Error(w, "Failed to retrieve permits", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(permits)
}

func ApprovePermit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	approver := r.URL.Query().Get("by")
	permit, err := usecase.ApprovePermit(id, approver)
	if err != nil {
		http.Error(w, "Approval failed", http.StatusInternalServerError)
		return
	}
	metrics.PermitGridApprovals.Inc()

	json.NewEncoder(w).Encode(permit)
}

func RejectPermit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	approver := r.URL.Query().Get("by")
	permit, err := usecase.RejectPermit(id, approver)
	if err != nil {
		http.Error(w, "Rejection failed", http.StatusInternalServerError)
		return
	}
	metrics.PermitGridRejections.Inc()

	json.NewEncoder(w).Encode(permit)
}
