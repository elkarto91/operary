package usecase

import (
	"github.com/elkarto91/operary/internal/corepad/model"
	"github.com/elkarto91/operary/internal/corepad/repo"
)

func CreateNote(content, author string, tags []string) (model.Note, error) {
	note := model.Note{
		Content: content,
		Author:  author,
		Tags:    tags,
	}
	return repo.SaveNote(note)
}

func GetNote(id string) (*model.Note, error) {
	return repo.GetNoteByID(id)
}
