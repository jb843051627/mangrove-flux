package model

import "time"

type FluxReading struct {
	ID           string       `json:"id"`
	DeploymentID string       `json:"deployment_id"`
	StationID    string       `json:"station_id"`
	ChamberID    string       `json:"chamber_id"`
	SensorKind   string       `json:"sensor_kind"`
	SampledAt    time.Time    `json:"sampled_at"`
	CO2Flux      float64      `json:"co2_flux"`
	CH4Flux      float64      `json:"ch4_flux"`
	TemperatureC float64      `json:"temperature_c"`
	SalinityPSU  float64      `json:"salinity_psu"`
	TideStage    string       `json:"tide_stage"`
	Quality      QualityState `json:"quality"`
	Source       string       `json:"source"`
	Checksum     uint32       `json:"checksum"`
	Notes        string       `json:"notes"`
}

func (r FluxReading) Validate() error {
	if r.ID == "" || r.DeploymentID == "" || r.StationID == "" || r.ChamberID == "" {
		return ErrInvalid
	}
	if r.SampledAt.IsZero() || r.SensorKind == "" {
		return ErrInvalid
	}
	if r.TemperatureC < -20 || r.TemperatureC > 70 || r.SalinityPSU < 0 || r.SalinityPSU > 80 {
		return ErrInvalid
	}
	return nil
}

func (r FluxReading) CarbonEquivalent() float64 { return r.CO2Flux + r.CH4Flux*28 }
