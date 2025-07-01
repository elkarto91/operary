package usecase

import (
	"context"

	"github.com/elkarto91/operary/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Result struct {
	Collection string `json:"collection"`
	Document   any    `json:"document"`
}

func Search(q, scope string) ([]Result, error) {
	db := config.GetMongoDB()
	collections := map[string]string{
		"corepad": "corepad_notes",
		"audit":   "audit_entries",
		"equip":   "equipment_ledger",
	}
	var results []Result
	for key, collName := range collections {
		if scope != "" && scope != key {
			continue
		}
		coll := db.Collection(collName)
		filter := bson.M{"$text": bson.M{"$search": q}}
		cursor, err := coll.Find(context.Background(), filter)
		if err != nil {
			continue
		}
		var docs []bson.M
		if err := cursor.All(context.Background(), &docs); err != nil {
			continue
		}
		for _, d := range docs {
			results = append(results, Result{Collection: key, Document: d})
		}
	}
	return results, nil
}
