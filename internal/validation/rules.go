package validation

import (
	"fmt"
	"time"
)

func SampleWindow(sampled, started time.Time) error {
	if sampled.Before(started.Add(-15 * time.Minute)) {
		return fmt.Errorf("sample predates deployment")
	}
	if sampled.After(started.Add(24 * time.Hour)) {
		return fmt.Errorf("sample outside deployment window")
	}
	return nil
}

func SalinityBand(value float64) string {
	switch {
	case value < 5:
		return "fresh"
	case value < 25:
		return "brackish"
	default:
		return "marine"
	}
}
