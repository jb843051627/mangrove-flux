package service

import (
	"context"
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
)

func (l *Lab) EvaluateReading(ctx context.Context, id string) (*model.FluxReading, error) {
	reading, err := l.readings.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	reading.Quality = policy.QualityFor(*reading, policy.DefaultLimits())
	if code, severity, message := policy.AlertFor(*reading); code != "" {
		alert := model.QualityAlert{ID: reading.ID + ":" + code, ReadingID: reading.ID, StationID: reading.StationID, Code: code, Severity: severity, State: model.AlertOpen, Message: message, CreatedAt: l.clock.Now()}
		if err := l.alerts.Save(ctx, &alert); err != nil {
			return nil, err
		}
	}
	if err := l.readings.Save(ctx, reading); err != nil {
		return nil, err
	}
	return reading, nil
}
func (l *Lab) QualitySweep(ctx context.Context, stationID string) (int, error) {
	if err := cancelled(ctx); err != nil {
		return 0, err
	}
	readings, err := l.readings.ListByStation(ctx, stationID)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, reading := range readings {
		if err := cancelled(ctx); err != nil {
			return count, err
		}
		if _, err := l.EvaluateReading(ctx, reading.ID); err != nil {
			return count, fmt.Errorf("quality sweep: %w", err)
		}
		count++
	}
	return count, nil
}
func (l *Lab) RecalculateQuality(ctx context.Context, reading model.FluxReading) model.QualityState {
	return policy.QualityFor(reading, policy.DefaultLimits())
}
