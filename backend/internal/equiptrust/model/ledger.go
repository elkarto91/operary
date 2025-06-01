package model

import "time"

type Status struct {
	System     string    `json:"system"`
	State      string    `json:"state"` // e.g., "OK", "WARN", "FAIL"
	LastUpdate time.Time `json:"last_update"`
	Notes      string    `json:"notes"`
}
