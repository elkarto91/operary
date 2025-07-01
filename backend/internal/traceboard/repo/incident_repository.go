package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/traceboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// InitMongoCollection stores the collection handle for incident reports.
func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("traceboard_incidents")
}

// Save inserts a new incident report with a creation timestamp.
func Save(report model.IncidentReport) (model.IncidentReport, error) {
	report.ID = primitive.NewObjectID()
	report.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.Background(), report)
	return report, err
}

// List returns all stored incident reports.
func List() ([]model.IncidentReport, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var results []model.IncidentReport
	err = cursor.All(context.Background(), &results)
	return results, err
}

// GetByID fetches a single report by its ObjectID.
func GetByID(id primitive.ObjectID) (model.IncidentReport, error) {
	var result model.IncidentReport
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	return result, err
}

// DeleteByID removes a report, used by admin tools.
func DeleteByID(id primitive.ObjectID) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
