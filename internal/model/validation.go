package model

import "math"

func Finite(value float64) bool { return !math.IsNaN(value) && !math.IsInf(value, 0) }

func ValidFlux(value float64) bool { return Finite(value) && value >= -5000 && value <= 5000 }

func ValidSensorKind(value string) bool {
	switch value {
	case "co2", "ch4", "temperature", "salinity":
		return true
	default:
		return false
	}
}
