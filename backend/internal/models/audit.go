package models

import (
	"context"
	"time"

	"github.com/elkarto91/operary/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuditLog struct {
	ID         string    `bson:"_id"`
	EntityType string    `bson:"entity_type"`
	EntityID   string    `bson:"entity_id"`
	Action     string    `bson:"action"`
	UserID     string    `bson:"user_id"`
	Timestamp  time.Time `bson:"timestamp"`
	Snapshot   any       `bson:"snapshot,omitempty"`
}

var auditCollection *mongo.Collection = config.OperaryDB.Collection("audit_logs")

func RecordAudit(entityType, entityID, action, userID string, snapshot any) error {
	entry := AuditLog{
		ID:         uuid.New().String(),
		EntityType: entityType,
		EntityID:   entityID,
		Action:     action,
		UserID:     userID,
		Timestamp:  time.Now(),
		Snapshot:   snapshot,
	}

	_, err := auditCollection.InsertOne(context.TODO(), entry)
	return err
}

func GetAuditLogs(entityType, entityID string) ([]AuditLog, error) {
	filter := bson.M{}
	if entityType != "" {
		filter["entity_type"] = entityType
	}
	if entityID != "" {
		filter["entity_id"] = entityID
	}

	cursor, err := auditCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var logs []AuditLog
	for cursor.Next(context.TODO()) {
		var entry AuditLog
		if err := cursor.Decode(&entry); err != nil {
			return nil, err
		}
		logs = append(logs, entry)
	}

	return logs, nil
}
