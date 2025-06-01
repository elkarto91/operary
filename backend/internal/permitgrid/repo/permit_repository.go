package repo

import (
	"context"
	"time"

	"github.com/elkarto91/operary/internal/permitgrid/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("permit_requests")
}

func SavePermit(p model.Permit) (model.Permit, error) {
	p.ID = primitive.NewObjectID()
	p.RequestedAt = time.Now()
	p.Status = "PENDING"
	_, err := collection.InsertOne(context.Background(), p)
	return p, err
}

func ListPermits() ([]model.Permit, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var results []model.Permit
	err = cursor.All(context.Background(), &results)
	return results, err
}

func UpdatePermitStatus(id primitive.ObjectID, status string, approver string) (model.Permit, error) {
	var permit model.Permit
	now := time.Now()

	update := bson.M{
		"$set": bson.M{
			"status":      status,
			"approved_by": approver,
			"approved_at": now,
		},
	}

	err := collection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, update).Decode(&permit)
	return permit, err
}
