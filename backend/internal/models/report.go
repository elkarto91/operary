package models

import (
	"fmt"
	"time"
)

// GenerateShiftReport creates a simple text summary of a shift and related tasks.
func GenerateShiftReport(shiftID string) (string, error) {
	shift, err := GetShiftByID(shiftID)
	if err != nil {
		return "", err
	}

	tasks, err := GetTasksByTimeRange(shift.StartedAt, shift.EndedAt)
	if err != nil {
		return "", err
	}

	report := fmt.Sprintf("Shift %s\nSupervisor: %s\nStart: %s\nEnd: %s\n", shift.ID, shift.SupervisorID, shift.StartedAt.Format(time.RFC3339), shift.EndedAt.Format(time.RFC3339))
	report += "Tasks:\n"
	for _, t := range tasks {
		report += fmt.Sprintf("- %s (%s)\n", t.Title, t.Status)
	}
	return report, nil
}
