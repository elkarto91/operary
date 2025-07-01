package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/corepad/usecase"
	"github.com/gorilla/mux"
)

// Request body structure for creating notes
type NoteRequest struct {
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
	Media   []string `json:"media"`
}

// POST /corepad/notes
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req NoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	note, err := usecase.CreateNote(req.Content, req.Author, req.Tags, req.Media)
	if err != nil {
		http.Error(w, "Could not create note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// GET /corepad/notes/{id}
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	note, err := usecase.GetNote(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}
