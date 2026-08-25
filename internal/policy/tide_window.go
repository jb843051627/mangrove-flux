package policy

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"time"
)

func WindowWeight(stage string) float64 {
	switch stage {
	case "high":
		return 1.20
	case "flood":
		return 1.10
	case "slack":
		return 1.00
	case "ebb":
		return 0.95
	case "low":
		return 0.85
	default:
		return 0.70
	}
}

func IsSuitableForDeployment(stage string, when time.Time) bool {
	if when.IsZero() || !model.Finite(WindowWeight(stage)) {
		return false
	}
	return stage != "storm" && stage != "unknown"
}

func ExposureHours(stage string, duration time.Duration) float64 {
	return duration.Hours() * WindowWeight(stage)
}
