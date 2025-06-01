package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"operary/internal/auditsync/model"
)

var auditors = []string{"Rajesh", "Lata", "Mohan", "Anita"}
var findings = []string{"No issues", "Minor deviation", "Check valve alignment", "PPE non-compliance"}
var tags = [][]string{
	{"safety", "routine"},
	{"quality", "followup"},
	{"esg", "incident"},
	{"equipment", "critical"},
}

func randomAudit() model.AuditEntry {
	return model.AuditEntry{
		ID:          primitive.NewObjectID(),
		ChecklistID: "CHK-" + primitive.NewObjectID().Hex()[0:6],
		Auditor:     auditors[rand.Intn(len(auditors))],
		Timestamp:   time.Now(),
		Passed:      rand.Intn(2) == 0,
		Findings:    findings[rand.Intn(len(findings))],
		Tags:        tags[rand.Intn(len(tags))],
		ShiftID:     "SHIFT-" + primitive.NewObjectID().Hex()[0:5],
	}
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("operary_dev")
	col := db.Collection("audit_entries")

	for i := 0; i < 10; i++ {
		audit := randomAudit()
		_, err := col.InsertOne(context.TODO(), audit)
		if err != nil {
			log.Println("Insert failed:", err)
		} else {
			log.Println("Inserted:", audit.Auditor, audit.Findings)
		}
	}
}
