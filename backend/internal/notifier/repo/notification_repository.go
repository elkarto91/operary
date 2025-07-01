package repo

import (
	"context"

	"github.com/elkarto91/operary/internal/notifier/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// InitMongoCollection sets the collection used for notifications.
func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("notifications")
	idx := mongo.IndexModel{Keys: bson.D{{Key: "timestamp", Value: -1}}}
	collection.Indexes().CreateOne(context.Background(), idx)
}

// Insert stores a new notification document.
func Insert(n model.Notification) (model.Notification, error) {
	_, err := collection.InsertOne(context.Background(), n)
	return n, err
}

// ListByUser retrieves recent notifications for a user.
func ListByUser(userID string, limit int) ([]model.Notification, error) {
	if limit <= 0 {
		limit = 50
	}
	opts := options.Find().SetSort(bson.M{"timestamp": -1}).SetLimit(int64(limit))
	cur, err := collection.Find(context.Background(), bson.M{"user_id": userID}, opts)
	if err != nil {
		return nil, err
	}
	var res []model.Notification
	if err := cur.All(context.Background(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
