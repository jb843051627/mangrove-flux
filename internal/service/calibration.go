package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
	"time"
)

func (l *Lab) CreateCalibration(ctx context.Context, profile model.CalibrationProfile) error {
	if profile.ID == "" || profile.SensorKind == "" || profile.StationID == "" || profile.Scale == 0 {
		return model.ErrInvalid
	}
	if profile.ValidUntil.IsZero() {
		return model.ErrInvalid
	}
	if _, err := l.requireStation(ctx, profile.StationID); err != nil {
		return err
	}
	profile.Points = model.ClonePoints(profile.Points)
	return l.calibrations.Save(ctx, &profile)
}
func (l *Lab) GetCalibration(ctx context.Context, id string) (*model.CalibrationProfile, error) {
	return l.calibrations.Get(ctx, id)
}
func (l *Lab) ActiveCalibration(ctx context.Context, sensorKind, stationID string, at time.Time) (*model.CalibrationProfile, error) {
	items, err := l.calibrations.ListBySensor(ctx, sensorKind, stationID)
	if err != nil {
		return nil, err
	}
	for i := range items {
		if items[i].UsableAt(at) {
			return &items[i], nil
		}
	}
	return nil, nil
}
func (l *Lab) ApplyCalibration(ctx context.Context, reading model.FluxReading, profile model.CalibrationProfile) (model.FluxReading, error) {
	if len(profile.Points) == 0 {
		return reading, model.ErrCalibration
	}
	value, err := policy.ApplyCalibration(reading.CO2Flux, profile)
	if err != nil {
		return reading, err
	}
	reading.CO2Flux = value
	return reading, nil
}
func (l *Lab) CalibrationDrift(ctx context.Context, profileID string) (float64, error) {
	profile, err := l.calibrations.Get(ctx, profileID)
	if err != nil {
		return 0, err
	}
	return policy.CalibrationDrift(*profile), nil
}
