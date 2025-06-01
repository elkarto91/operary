package main

import (
    "context"
    "log"
    "time"
    "math/rand"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "operary/internal/equiptrust/model"
)

var users = []string{"Anil", "Priya", "Sameer", "Kavya"}
var purposes = []string{"Routine inspection", "Emergency repair", "Daily usage", "Quality check"}
var equipments = []string{"DRILL-001", "METER-203", "WRENCH-88", "CAMERA-72"}

func randomLedgerEntry() model.EquipmentLedger {
    return model.EquipmentLedger{
        ID:           primitive.NewObjectID(),
        EquipmentID:  equipments[rand.Intn(len(equipments))],
        UsedBy:       users[rand.Intn(len(users))],
        Purpose:      purposes[rand.Intn(len(purposes))],
        ShiftID:      "SHIFT-" + primitive.NewObjectID().Hex()[0:6],
        CheckedOutAt: time.Now().Add(-time.Duration(rand.Intn(72)) * time.Hour),
        Notes:        "Simulated seed entry",
    }
}

func main() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    col := client.Database("operary_dev").Collection("equipment_ledger")

    for i := 0; i < 10; i++ {
        entry := randomLedgerEntry()
        _, err := col.InsertOne(context.TODO(), entry)
        if err != nil {
            log.Println("Insert failed:", err)
        } else {
            log.Println("Inserted:", entry.EquipmentID)
        }
    }
}