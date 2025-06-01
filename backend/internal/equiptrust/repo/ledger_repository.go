package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/opsmirror/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("opsmirror_status")
}

func GetAllStatuses() ([]model.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var results []model.Status
	err = cursor.All(ctx, &results)
	return results, err
}

func UpsertStatus(status model.Status) error {
	filter := bson.M{"system": status.System}
	update := bson.M{"$set": status}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}
