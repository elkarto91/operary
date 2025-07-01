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

func GetByID(id primitive.ObjectID) (model.SensorEvent, error) {
	var result model.SensorEvent
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	return result, err
}

func DeleteByID(id primitive.ObjectID) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
