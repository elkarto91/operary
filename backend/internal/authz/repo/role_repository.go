package repo

import (
	"context"

	"github.com/elkarto91/operary/internal/authz/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) {
	collection = db.Collection("user_roles")
}

func UpsertRole(userID, role, orgID string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": bson.M{"role": role, "org_id": orgID}}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func GetRole(userID string) (model.UserRole, error) {
	var ur model.UserRole
	err := collection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&ur)
	return ur, err
}
