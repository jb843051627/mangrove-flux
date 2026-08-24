package model

import "time"

type FluxReport struct {
	ID           string    `json:"id"`
	StationID    string    `json:"station_id"`
	Day          string    `json:"day"`
	Samples      int       `json:"samples"`
	GoodSamples  int       `json:"good_samples"`
	MeanCO2      float64   `json:"mean_co2"`
	MeanCH4      float64   `json:"mean_ch4"`
	NetCarbon    float64   `json:"net_carbon"`
	TideExposure float64   `json:"tide_exposure"`
	AlertCount   int       `json:"alert_count"`
	GeneratedAt  time.Time `json:"generated_at"`
}
