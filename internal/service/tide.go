package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/validation"
	"sort"
)

func (l *Lab) RecordTide(ctx context.Context, observation model.TideObservation) error {
	if err := observation.Validate(); err != nil {
		return err
	}
	if _, err := l.requireStation(ctx, observation.StationID); err != nil {
		return err
	}
	if !validation.TideCompatible(observation.Stage) {
		return model.ErrInvalid
	}
	return l.tides.Save(ctx, &observation)
}
func (l *Lab) ListTides(ctx context.Context, stationID string) ([]model.TideObservation, error) {
	return l.tides.ListByStation(ctx, stationID)
}
func (l *Lab) LatestTide(ctx context.Context, stationID string) (*model.TideObservation, error) {
	values, err := l.tides.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, model.ErrNotFound
	}
	sort.Slice(values, func(i, j int) bool { return values[i].ObservedAt.After(values[j].ObservedAt) })
	return &values[0], nil
}
func (l *Lab) TideHeight(ctx context.Context, stationID string) (float64, error) {
	tide, err := l.LatestTide(ctx, stationID)
	if err != nil {
		return 0, err
	}
	return tide.HeightM, nil
}
func (l *Lab) TideStage(ctx context.Context, stationID string) (string, error) {
	tide, err := l.LatestTide(ctx, stationID)
	if err != nil {
		return "unknown", err
	}
	return tide.Stage, nil
}
