package usecase

import (
	"github.com/elkarto91/operary/internal/opsmirror/model"
	"github.com/elkarto91/operary/internal/opsmirror/repo"
)

func FetchStatuses() ([]model.Status, error) {
	return repo.GetAllStatuses()
}
