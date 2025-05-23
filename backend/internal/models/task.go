
package models

import (
    "context"
    "time"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "operary/config"
)

type Task struct {
    ID          string             `bson:"_id"`
    Title       string             `bson:"title"`
    AssignedTo  string             `bson:"assigned_to"`
    Status      string             `bson:"status"`
    CreatedAt   time.Time          `bson:"created_at"`
    Notes       []string           `bson:"notes,omitempty"`
}

var taskCollection *mongo.Collection = config.OperaryDB.Collection("tasks")

func CreateTask(title, assignee string) (*Task, error) {
    task := Task{
        ID:         uuid.New().String(),
        Title:      title,
        AssignedTo: assignee,
        Status:     "open",
        CreatedAt:  time.Now(),
    }

    _, err := taskCollection.InsertOne(context.TODO(), task)
    if err != nil {
        return nil, err
    }

    return &task, nil
}

func GetAllTasks() ([]Task, error) {
    var tasks []Task
    cursor, err := taskCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var task Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}
