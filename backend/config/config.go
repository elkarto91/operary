package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var OperaryDB *mongo.Database

func InitMongo() {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	if uri == "" || dbName == "" {
		log.Fatal("MONGO_URI and MONGO_DB must be set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	MongoClient = client
	OperaryDB = client.Database(dbName)

	fmt.Println("✅ MongoDB connected:", dbName)
}
func GetMongoDB() *mongo.Database {
	// Return the initialized database instance used by the rest of the
	// application. This was previously returning an unassigned variable
	// which caused nil pointer errors when accessing the DB from other
	// packages.
	// During initialization OperaryDB is set to the active database
	// connection. Simply return it here so other packages can interact
	// with MongoDB after initialization.
	return OperaryDB
}
