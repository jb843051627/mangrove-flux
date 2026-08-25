package policy

import (
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type FluxLimits struct{ CO2Min, CO2Max, CH4Min, CH4Max float64 }

func DefaultLimits() FluxLimits {
	return FluxLimits{CO2Min: -1500, CO2Max: 3000, CH4Min: -100, CH4Max: 500}
}

func ValidateFlux(reading model.FluxReading, limits FluxLimits) error {
	if !model.ValidFlux(reading.CO2Flux) || reading.CO2Flux < limits.CO2Min || reading.CO2Flux > limits.CO2Max {
		return fmt.Errorf("%w: co2 flux", model.ErrQuality)
	}
	if !model.ValidFlux(reading.CH4Flux) || reading.CH4Flux < limits.CH4Min || reading.CH4Flux > limits.CH4Max {
		return fmt.Errorf("%w: ch4 flux", model.ErrQuality)
	}
	return nil
}

func QualityFor(reading model.FluxReading, limits FluxLimits) model.QualityState {
	if ValidateFlux(reading, limits) != nil {
		return model.QualityRejected
	}
	if reading.TemperatureC < -5 || reading.TemperatureC > 55 || reading.SalinityPSU > 60 {
		return model.QualityReview
	}
	return model.QualityGood
}

func CarbonWeight(reading model.FluxReading, tideStage string) float64 {
	weight := 1.0
	if tideStage == "high" || tideStage == "flood" {
		weight = 1.15
	}
	if tideStage == "low" || tideStage == "ebb" {
		weight = 0.90
	}
	return reading.CarbonEquivalent() * weight
}
