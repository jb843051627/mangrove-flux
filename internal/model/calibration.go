package model

import "time"

type CalibrationPoint struct {
	Input  float64 `json:"input"`
	Output float64 `json:"output"`
}

type CalibrationProfile struct {
	ID         string             `json:"id"`
	SensorKind string             `json:"sensor_kind"`
	StationID  string             `json:"station_id"`
	Offset     float64            `json:"offset"`
	Scale      float64            `json:"scale"`
	ValidFrom  time.Time          `json:"valid_from"`
	ValidUntil time.Time          `json:"valid_until"`
	Approved   bool               `json:"approved"`
	Points     []CalibrationPoint `json:"points"`
}

func (p CalibrationProfile) UsableAt(now time.Time) bool {
	return p.Approved && p.Scale != 0 && !now.Before(p.ValidFrom) && now.Before(p.ValidUntil)
}
