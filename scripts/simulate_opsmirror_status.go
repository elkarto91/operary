package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/elkarto91/operary/backend/internal/opsmirror/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var systems = []string{"Boiler A", "Conveyor 3", "Packaging Unit", "Cooling System"}

var states = []string{"OK", "WARN", "FAIL"}

func randomStatus(system string) model.Status {
	return model.Status{
		System:     system,
		State:      states[rand.Intn(len(states))],
		LastUpdate: time.Now(),
		Notes:      "Simulated by bot",
	}
}

func main() {
	client, err := mongo.Connect(nil, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}

	db := client.Database("operary_dev")
	col := db.Collection("opsmirror_status")

	log.Println("📡 Starting OpsMirror Simulator...")

	for {
		for _, system := range systems {
			status := randomStatus(system)
			_, err := col.UpdateOne(
				nil,
				bson.M{"system": status.System},
				bson.M{"$set": status},
				options.Update().SetUpsert(true),
			)
			if err != nil {
				log.Println("⚠️ Update error:", err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
