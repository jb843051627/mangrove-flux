package validation

import "time"

func InWindow(value, start, end time.Time) bool { return !value.Before(start) && !value.After(end) }
func Recent(value time.Time, limit time.Duration, now time.Time) bool {
	return !value.After(now) && now.Sub(value) <= limit
}
func TideCompatible(stage string) bool {
	switch stage {
	case "ebb", "flood", "slack", "high", "low":
		return true
	default:
		return false
	}
}
