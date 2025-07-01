package usecase

import "strings"

// DetectPriority assigns a basic priority score based on keywords.
func DetectPriority(content string) float64 {
	c := strings.ToLower(content)
	switch {
	case strings.Contains(c, "urgent"):
		return 0.9
	case strings.Contains(c, "low"):
		return 0.2
	default:
		return 0.5
	}
}

// ShouldEscalate recommends escalation based on simple rules.
func ShouldEscalate(content string) bool {
	c := strings.ToLower(content)
	return strings.Contains(c, "escalate") || strings.Contains(c, "critical")
}
