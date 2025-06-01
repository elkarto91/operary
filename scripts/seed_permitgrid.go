package main

import (
    "context"
    "log"
    "time"
    "math/rand"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "operary/internal/permitgrid/model"
)

var requesters = []string{"Meera", "Vinay", "Akash", "Devika"}
var titles = []string{"Confined Space Entry", "Hot Work Permit", "Electrical Isolation", "Working at Height"}
var descriptions = []string{
    "Entering vessel P-101 for inspection",
    "Welding work near tank T-23",
    "De-energizing MCC panel A",
    "Installing safety nets on tower",
}
var tags = [][]string{
    {"safety", "urgent"},
    {"maintenance"},
    {"electrical", "critical"},
    {"height", "precaution"},
}

func randomPermit() model.Permit {
    return model.Permit{
        ID:          primitive.NewObjectID(),
        Title:       titles[rand.Intn(len(titles))],
        Description: descriptions[rand.Intn(len(descriptions))],
        RequestedBy: requesters[rand.Intn(len(requesters))],
        Status:      "PENDING",
        RequestedAt: time.Now(),
        ShiftID:     "SHIFT-" + primitive.NewObjectID().Hex()[0:6],
        Tags:        tags[rand.Intn(len(tags))],
    }
}

func main() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    col := client.Database("operary_dev").Collection("permit_requests")

    for i := 0; i < 10; i++ {
        permit := randomPermit()
        _, err := col.InsertOne(context.TODO(), permit)
        if err != nil {
            log.Println("Insert failed:", err)
        } else {
            log.Println("Inserted permit:", permit.Title)
        }
    }
}