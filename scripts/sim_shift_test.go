package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiBase = "http://localhost:8080/v1"
const orgToken = "your-org-token-here"

var client = &http.Client{}

func main() {
	// 1. Start shift
	shiftID := startShift("supervisor-123")

	// 2. Create tasks
	task1 := createTask("Check pump pressure", "operator-101", "")
	task2 := createTask("Inspect valve leak", "operator-102", "")

	// 3. Update task status
	updateTaskStatus(task1, "in_progress")
	time.Sleep(2 * time.Second)
	updateTaskStatus(task1, "closed")

	updateTaskStatus(task2, "escalated")
	time.Sleep(1 * time.Second)

	// 4. Close shift
	closeShift(shiftID)

	fmt.Println("✅ Simulated shift complete.")
}

func startShift(supervisorID string) string {
	body := map[string]string{"supervisor_id": supervisorID}
	res := post("/shifts", body)
	var resp map[string]interface{}
	json.Unmarshal(res, &resp)
	fmt.Println("🕓 Shift started:", resp["id"])
	return resp["id"].(string)
}

func createTask(title, assignee, due string) string {
	body := map[string]string{
		"title":       title,
		"assigned_to": assignee,
	}
	if due != "" {
		body["due_time"] = due
	}
	res := post("/tasks", body)
	var resp map[string]interface{}
	json.Unmarshal(res, &resp)
	fmt.Println("📋 Task created:", title)
	return resp["id"].(string)
}

func updateTaskStatus(id, status string) {
	body := map[string]string{"status": status}
	post(fmt.Sprintf("/tasks/%s", id), body)
	fmt.Println("🔄 Task", id, "→", status)
}

func closeShift(id string) {
	post(fmt.Sprintf("/shifts/%s/close", id), nil)
	fmt.Println("🔒 Shift", id, "closed.")
}

func post(path string, data interface{}) []byte {
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", apiBase+path, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Org-Token", orgToken)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}