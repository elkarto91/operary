package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/twinboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// InitMongoCollection stores the MongoDB collection handle.
func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("twin_snapshots")
}

// SaveSnapshot inserts a snapshot into MongoDB.
func SaveSnapshot(s model.TwinSnapshot) (model.TwinSnapshot, error) {
	if s.Timestamp.IsZero() {
		s.Timestamp = time.Now()
	}
	_, err := collection.InsertOne(context.Background(), s)
	return s, err
}

// GetLatest returns the most recent snapshot.
func GetLatest() (model.TwinSnapshot, error) {
	opts := options.FindOne().SetSort(bson.M{"timestamp": -1})
	var snap model.TwinSnapshot
	err := collection.FindOne(context.Background(), bson.M{}, opts).Decode(&snap)
	return snap, err
}

// ListRange returns snapshots between start and end timestamps.
func ListRange(start, end time.Time) ([]model.TwinSnapshot, error) {
	filter := bson.M{"timestamp": bson.M{"$gte": start, "$lte": end}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var results []model.TwinSnapshot
	err = cursor.All(context.Background(), &results)
	return results, err
}
