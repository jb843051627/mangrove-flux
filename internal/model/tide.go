package model

import "time"

type TideObservation struct {
	ID         string    `json:"id"`
	StationID  string    `json:"station_id"`
	ObservedAt time.Time `json:"observed_at"`
	HeightM    float64   `json:"height_m"`
	Stage      string    `json:"stage"`
}

func (t TideObservation) Validate() error {
	if t.ID == "" || t.StationID == "" || t.ObservedAt.IsZero() {
		return ErrInvalid
	}
	if t.HeightM < -5 || t.HeightM > 15 {
		return ErrInvalid
	}
	return nil
}
