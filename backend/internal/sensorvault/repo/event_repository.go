package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/sensorvault/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("sensor_events")
}

func Save(event model.SensorEvent) (model.SensorEvent, error) {
	event.ID = primitive.NewObjectID()
	event.Timestamp = time.Now()
	_, err := collection.InsertOne(context.Background(), event)
	return event, err
}

func List() ([]model.SensorEvent, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var results []model.SensorEvent
	err = cursor.All(context.Background(), &results)
	return results, err
}
