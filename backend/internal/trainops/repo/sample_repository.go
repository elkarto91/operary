package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/trainops/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("training_samples")
}

func Save(sample model.TrainingSample) (model.TrainingSample, error) {
	sample.ID = primitive.NewObjectID()
	sample.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.Background(), sample)
	return sample, err
}

func List() ([]model.TrainingSample, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var results []model.TrainingSample
	err = cursor.All(context.Background(), &results)
	return results, err
}

func UpdateFeedback(id primitive.ObjectID, fb string) (model.TrainingSample, error) {
	var updated model.TrainingSample
	update := bson.M{"$set": bson.M{"feedback": fb}}
	err := collection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, update).Decode(&updated)
	return updated, err
}
