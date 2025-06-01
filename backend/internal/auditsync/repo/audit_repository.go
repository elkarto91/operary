package repo

import (
	"context"

	"github.com/elkarto91/operary/internal/auditsync/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("audit_entries")
}

func SaveAudit(a model.AuditEntry) (model.AuditEntry, error) {
	a.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), a)
	return a, err
}

func ListAudits() ([]model.AuditEntry, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var results []model.AuditEntry
	err = cursor.All(context.Background(), &results)
	return results, err
}
