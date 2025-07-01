package models

import (
	"context"
	"time"

	"github.com/elkarto91/operary/config"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID         string    `bson:"_id"`
	Title      string    `bson:"title"`
	AssignedTo string    `bson:"assigned_to"`
	Status     string    `bson:"status"`
	CreatedAt  time.Time `bson:"created_at"`
	Notes      []string  `bson:"notes,omitempty"`
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

// GetTasksByTimeRange returns all tasks created between the given start and end
// times. Tasks are filtered based on the "created_at" timestamp stored in the
// database.
func GetTasksByTimeRange(start, end time.Time) ([]Task, error) {
	filter := bson.M{
		"created_at": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	cursor, err := taskCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []Task
	for cursor.Next(context.TODO()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdateTask updates the task status and/or appends a note. The updated task is
// returned after modification. If neither status nor note is provided, the task
// is left unchanged and the current record is returned.
func UpdateTask(id, status, note string) (*Task, error) {
	filter := bson.M{"_id": id}

	update := bson.M{}
	setFields := bson.M{}
	if status != "" {
		setFields["status"] = status
	}
	if len(setFields) > 0 {
		update["$set"] = setFields
	}
	if note != "" {
		update["$push"] = bson.M{"notes": note}
	}

	if len(update) == 0 {
		// Nothing to update; just return the existing task
		var existing Task
		err := taskCollection.FindOne(context.TODO(), filter).Decode(&existing)
		if err != nil {
			return nil, err
		}
		return &existing, nil
	}

	// Apply the update and then fetch the new record
	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var updated Task
	err = taskCollection.FindOne(context.TODO(), filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}
