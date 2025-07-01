package usecase

import (
	equip "github.com/elkarto91/operary/internal/equiptrust/usecase"
	"github.com/elkarto91/operary/internal/models"
	"github.com/elkarto91/operary/internal/opsmirror/usecase"
	"github.com/elkarto91/operary/internal/twinboard/model"
	"github.com/elkarto91/operary/internal/twinboard/repo"
	"time"
)

// GenerateSnapshot reads current system state and stores a TwinSnapshot.
func GenerateSnapshot() (model.TwinSnapshot, error) {
	tasks, _ := models.GetAllTasks()

	var activeTasks []string
	workerSet := make(map[string]struct{})
	for _, t := range tasks {
		activeTasks = append(activeTasks, t.ID)
		if t.AssignedTo != "" {
			workerSet[t.AssignedTo] = struct{}{}
		}
	}
	var activeWorkers []string
	for w := range workerSet {
		activeWorkers = append(activeWorkers, w)
	}

	ledger, _ := equip.ListLedger()
	equipment := make(map[string]string)
	for _, e := range ledger {
		if e.ReturnedAt == nil {
			equipment[e.EquipmentID] = e.Purpose
		}
	}

	statuses, _ := usecase.FetchStatuses()
	var warnings []string
	for _, s := range statuses {
		if s.State == "WARN" {
			warnings = append(warnings, s.System)
		}
	}
	score := 1.0
	if len(statuses) > 0 {
		score = 1.0 - float64(len(warnings))/float64(len(statuses))
	}

	snap := model.TwinSnapshot{
		Timestamp:     time.Now(),
		ActiveWorkers: activeWorkers,
		ActiveTasks:   activeTasks,
		EquipmentUsed: equipment,
		Warnings:      warnings,
		StatusScore:   score,
	}
	saved, err := repo.SaveSnapshot(snap)
	return saved, err
}

// GetLiveSnapshot returns the most recent TwinSnapshot.
func GetLiveSnapshot() (model.TwinSnapshot, error) {
	return repo.GetLatest()
}

// GetHistory returns snapshots between start and end time.
func GetHistory(start, end time.Time) ([]model.TwinSnapshot, error) {
	return repo.ListRange(start, end)
}
