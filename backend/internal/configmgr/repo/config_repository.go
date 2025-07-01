package repo

import (
	"context"

	"github.com/elkarto91/operary/internal/configmgr/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func InitMongoCollection(db *mongo.Database) { collection = db.Collection("configs") }

func SetConfig(entry model.ConfigEntry) error {
	filter := bson.M{"key": entry.Key, "org_id": entry.OrgID}
	update := bson.M{"$set": bson.M{"value": entry.Value}}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func GetConfig(key, orgID string) (model.ConfigEntry, error) {
	var c model.ConfigEntry
	err := collection.FindOne(context.Background(), bson.M{"key": key, "org_id": orgID}).Decode(&c)
	if err == mongo.ErrNoDocuments && orgID != "" {
		// fallback to global
		err = collection.FindOne(context.Background(), bson.M{"key": key, "org_id": ""}).Decode(&c)
	}
	return c, err
}
