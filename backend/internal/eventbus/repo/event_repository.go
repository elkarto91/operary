package repo

import (
	"context"

	"github.com/elkarto91/operary/internal/eventbus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("events")
	// ensure index on timestamp
	idx := mongo.IndexModel{Keys: bson.D{{Key: "timestamp", Value: -1}}}
	collection.Indexes().CreateOne(context.Background(), idx)
}

func InsertEvent(e model.Event) (model.Event, error) {
	_, err := collection.InsertOne(context.Background(), e)
	return e, err
}

func ListEvents(limit int, filter bson.M) ([]model.Event, error) {
	opts := options.Find().SetSort(bson.M{"timestamp": -1}).SetLimit(int64(limit))
	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	var res []model.Event
	if err := cur.All(context.Background(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
