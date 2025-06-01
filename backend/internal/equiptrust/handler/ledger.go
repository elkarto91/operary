package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/elkarto91/operary/internal/equiptrust/usecase"
	metrics "github.com/elkarto91/operary/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func CheckoutEquipment(w http.ResponseWriter, r *http.Request) {
	var req usecase.EquipmentCheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	entry, err := usecase.Checkout(req)
	if err != nil {
		http.Error(w, "Failed to checkout equipment", http.StatusInternalServerError)
		return
	}
	metrics.EquipTrustCheckouts.Inc()

	json.NewEncoder(w).Encode(entry)
}

func ReturnEquipment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	entry, err := usecase.Return(id, time.Now())
	if err != nil {
		http.Error(w, "Return failed", http.StatusInternalServerError)
		return
	}
	metrics.EquipTrustReturns.Inc()

	json.NewEncoder(w).Encode(entry)
}

func ListLedger(w http.ResponseWriter, r *http.Request) {
	entries, err := usecase.ListLedger()
	if err != nil {
		http.Error(w, "Failed to retrieve ledger", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entries)
}
