package model

import "time"

type WeatherObservation struct {
	ID              string    `json:"id"`
	StationID       string    `json:"station_id"`
	ObservedAt      time.Time `json:"observed_at"`
	RainMM          float64   `json:"rain_mm"`
	AirTemperatureC float64   `json:"air_temperature_c"`
	WindMPS         float64   `json:"wind_mps"`
	CloudFraction   float64   `json:"cloud_fraction"`
}

func (w WeatherObservation) Validate() error {
	if w.ID == "" || w.StationID == "" || w.ObservedAt.IsZero() {
		return ErrInvalid
	}
	if w.RainMM < 0 || w.WindMPS < 0 || w.CloudFraction < 0 || w.CloudFraction > 1 {
		return ErrInvalid
	}
	return nil
}
