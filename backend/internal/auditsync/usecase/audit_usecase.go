package usecase

import (
	"time"

	"github.com/elkarto91/operary/internal/auditsync/model"
	"github.com/elkarto91/operary/internal/auditsync/repo"
)

type AuditEntryRequest struct {
	ChecklistID string   `json:"checklist_id"`
	Auditor     string   `json:"auditor"`
	Passed      bool     `json:"passed"`
	Findings    string   `json:"findings"`
	Tags        []string `json:"tags"`
	ShiftID     string   `json:"shift_id"`
}

func CreateAudit(req AuditEntryRequest) (model.AuditEntry, error) {
	entry := model.AuditEntry{
		ChecklistID: req.ChecklistID,
		Auditor:     req.Auditor,
		Passed:      req.Passed,
		Findings:    req.Findings,
		Tags:        req.Tags,
		ShiftID:     req.ShiftID,
		Timestamp:   time.Now(),
	}
	return repo.SaveAudit(entry)
}

func ListAudits() ([]model.AuditEntry, error) {
	return repo.ListAudits()
}
