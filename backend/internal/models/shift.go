
package models

import (
    "context"
    "time"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "operary/config"
)

type Shift struct {
    ID           string    `bson:"_id"`
    SupervisorID string    `bson:"supervisor_id"`
    StartedAt    time.Time `bson:"started_at"`
    EndedAt      time.Time `bson:"ended_at,omitempty"`
    Closed       bool      `bson:"closed"`
}

var shiftCollection *mongo.Collection = config.OperaryDB.Collection("shifts")

func StartShift(supervisorID string) (*Shift, error) {
    shift := Shift{
        ID:           uuid.New().String(),
        SupervisorID: supervisorID,
        StartedAt:    time.Now(),
        Closed:       false,
    }

    _, err := shiftCollection.InsertOne(context.TODO(), shift)
    if err != nil {
        return nil, err
    }

    return &shift, nil
}

func CloseShift(shiftID string) error {
    filter := bson.M{"_id": shiftID}
    update := bson.M{
        "$set": bson.M{
            "closed":   true,
            "ended_at": time.Now(),
        },
    }

    _, err := shiftCollection.UpdateOne(context.TODO(), filter, update)
    return err
}
