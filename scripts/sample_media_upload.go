package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBase = "http://localhost:8080/v1"
const orgToken = "your-org-token-here"

func main() {
	note := map[string]interface{}{
		"content": "Pump status photo attached",
		"author":  "tester",
		"tags":    []string{"shift-a"},
		"media":   []string{"https://example.com/photo.jpg"},
	}
	data, _ := json.Marshal(note)
	req, _ := http.NewRequest("POST", apiBase+"/corepad/notes", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Org-Token", orgToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
