package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/supplymesh/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var suppliers *mongo.Collection
var deliveries *mongo.Collection
var escalations *mongo.Collection

func InitMongoCollections(db *mongo.Database) {
	suppliers = db.Collection("suppliers")
	deliveries = db.Collection("deliveries")
	escalations = db.Collection("escalations")
}

func SaveSupplier(s model.Supplier) (model.Supplier, error) {
	s.ID = primitive.NewObjectID()
	_, err := suppliers.InsertOne(context.Background(), s)
	return s, err
}

func ListSuppliers() ([]model.Supplier, error) {
	cursor, err := suppliers.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var results []model.Supplier
	err = cursor.All(context.Background(), &results)
	return results, err
}

func SaveDelivery(d model.Delivery) (model.Delivery, error) {
	d.ID = primitive.NewObjectID()
	d.Delivered = false
	_, err := deliveries.InsertOne(context.Background(), d)
	return d, err
}

func ListDeliveries() ([]model.Delivery, error) {
	cursor, err := deliveries.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var results []model.Delivery
	err = cursor.All(context.Background(), &results)
	return results, err
}

func SaveEscalation(e model.Escalation) (model.Escalation, error) {
	e.ID = primitive.NewObjectID()
	e.Timestamp = time.Now()
	_, err := escalations.InsertOne(context.Background(), e)
	return e, err
}
