package usecase

import (
	"time"

	"github.com/elkarto91/operary/internal/supplymesh/model"
	"github.com/elkarto91/operary/internal/supplymesh/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierRequest struct {
	Name string `json:"name"`
	SLA  int    `json:"sla"`
}

type DeliveryRequest struct {
	SupplierID   string `json:"supplier_id"`
	Item         string `json:"item"`
	ExpectedDate string `json:"expected_date"`
}

type EscalationRequest struct {
	DeliveryID string `json:"delivery_id"`
	Reason     string `json:"reason"`
}

func AddSupplier(req SupplierRequest) (model.Supplier, error) {
	supplier := model.Supplier{Name: req.Name, SLA: req.SLA}
	return repo.SaveSupplier(supplier)
}

func ListSuppliers() ([]model.Supplier, error) {
	return repo.ListSuppliers()
}

func AddDelivery(req DeliveryRequest) (model.Delivery, error) {
	supID, err := primitive.ObjectIDFromHex(req.SupplierID)
	if err != nil {
		return model.Delivery{}, err
	}
	expected, err := time.Parse(time.RFC3339, req.ExpectedDate)
	if err != nil {
		return model.Delivery{}, err
	}
	d := model.Delivery{SupplierID: supID, Item: req.Item, ExpectedDate: expected}
	return repo.SaveDelivery(d)
}

func ListDeliveries() ([]model.Delivery, error) {
	return repo.ListDeliveries()
}

func Escalate(req EscalationRequest) (model.Escalation, error) {
	dID, err := primitive.ObjectIDFromHex(req.DeliveryID)
	if err != nil {
		return model.Escalation{}, err
	}
	e := model.Escalation{DeliveryID: dID, Reason: req.Reason}
	return repo.SaveEscalation(e)
}
