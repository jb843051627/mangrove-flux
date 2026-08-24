package model

import "time"

type Chamber struct {
	ID            string    `json:"id"`
	StationID     string    `json:"station_id"`
	Serial        string    `json:"serial"`
	VolumeL       float64   `json:"volume_l"`
	SensorKind    string    `json:"sensor_kind"`
	CalibrationID string    `json:"calibration_id"`
	InstalledAt   time.Time `json:"installed_at"`
	Active        bool      `json:"active"`
}

func (c Chamber) Validate() error {
	if c.ID == "" || c.StationID == "" || c.Serial == "" || c.SensorKind == "" {
		return ErrInvalid
	}
	if c.VolumeL <= 0 {
		return ErrInvalid
	}
	return nil
}
