package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

func (l *Lab) RecordWeather(ctx context.Context, observation model.WeatherObservation) error {
	if err := observation.Validate(); err != nil {
		return err
	}
	if _, err := l.requireStation(ctx, observation.StationID); err != nil {
		return err
	}
	return l.weather.Save(ctx, &observation)
}
func (l *Lab) ListWeather(ctx context.Context, stationID string) ([]model.WeatherObservation, error) {
	return l.weather.ListByStation(ctx, stationID)
}
func (l *Lab) RainRisk(ctx context.Context, stationID string) (string, error) {
	values, err := l.weather.ListByStation(ctx, stationID)
	if err != nil {
		return "unknown", err
	}
	if len(values) == 0 {
		return "unknown", nil
	}
	latest := values[len(values)-1]
	switch {
	case latest.RainMM >= 30:
		return "critical", nil
	case latest.RainMM >= 10:
		return "watch", nil
	default:
		return "normal", nil
	}
}
