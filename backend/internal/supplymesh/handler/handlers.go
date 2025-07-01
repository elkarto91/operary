package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/supplymesh/usecase"
)

func AddSupplier(w http.ResponseWriter, r *http.Request) {
	var req usecase.SupplierRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	supplier, err := usecase.AddSupplier(req)
	if err != nil {
		http.Error(w, "Failed to save supplier", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(supplier)
}

func ListSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := usecase.ListSuppliers()
	if err != nil {
		http.Error(w, "Failed to retrieve suppliers", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(suppliers)
}

func AddDelivery(w http.ResponseWriter, r *http.Request) {
	var req usecase.DeliveryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	delivery, err := usecase.AddDelivery(req)
	if err != nil {
		http.Error(w, "Failed to save delivery", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(delivery)
}

func ListDeliveries(w http.ResponseWriter, r *http.Request) {
	deliveries, err := usecase.ListDeliveries()
	if err != nil {
		http.Error(w, "Failed to retrieve deliveries", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(deliveries)
}

func EscalateIssue(w http.ResponseWriter, r *http.Request) {
	var req usecase.EscalationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	escalation, err := usecase.Escalate(req)
	if err != nil {
		http.Error(w, "Failed to escalate issue", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(escalation)
}
