package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/corepad/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("corepad_notes")
}

func SaveNote(n model.Note) (model.Note, error) {
	n.ID = primitive.NewObjectID()
	n.Timestamp = time.Now()
	_, err := collection.InsertOne(context.Background(), n)
	return n, err
}

func GetNoteByID(id string) (*model.Note, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result model.Note
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
