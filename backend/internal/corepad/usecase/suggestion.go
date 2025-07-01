package usecase

import (
	"strings"
)

// DetectPriority assigns a fake priority score based on keyword presence
func DetectPriority(content string) float64 {
	content = strings.ToLower(content)
	if strings.Contains(content, "urgent") || strings.Contains(content, "immediate") {
		return 0.9
	}
	if strings.Contains(content, "warning") {
		return 0.7
	}
	return 0.3
}

// ShouldEscalate returns true if content matches fake escalation keywords
func ShouldEscalate(content string) bool {
	content = strings.ToLower(content)
	keywords := []string{"fail", "accident", "shutdown", "hazard", "breach"}
	for _, word := range keywords {
		if strings.Contains(content, word) {
			return true
		}
	}
	return false
}
