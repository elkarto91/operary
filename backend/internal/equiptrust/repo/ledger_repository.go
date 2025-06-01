package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/equiptrust/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("equipment_ledger")
}

func Save(entry model.EquipmentLedger) (model.EquipmentLedger, error) {
	entry.ID = primitive.NewObjectID()
	entry.CheckedOutAt = time.Now()
	_, err := collection.InsertOne(context.Background(), entry)
	return entry, err
}

func UpdateReturn(id primitive.ObjectID, returnTime time.Time) (model.EquipmentLedger, error) {
	var entry model.EquipmentLedger
	update := bson.M{
		"$set": bson.M{
			"returned_at": returnTime,
		},
	}

	err := collection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, update).Decode(&entry)
	return entry, err
}

func List() ([]model.EquipmentLedger, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var results []model.EquipmentLedger
	err = cursor.All(context.Background(), &results)
	return results, err
}
