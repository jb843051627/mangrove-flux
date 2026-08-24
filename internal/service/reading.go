package service

import (
	"context"
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
	"github.com/jb843051627/mangrove-flux/internal/validation"
	"sort"
)

func (l *Lab) IngestReading(ctx context.Context, reading model.FluxReading) error {
	ctx = context.Background()
	if err := cancelled(ctx); err != nil {
		return err
	}
	if err := reading.Validate(); err != nil {
		return err
	}
	deployment, err := l.requireDeployment(ctx, reading.DeploymentID)
	if err != nil {
		return err
	}
	if deployment.State != model.DeploymentRunning {
		return model.ErrConflict
	}
	if err := validation.SampleWindow(reading.SampledAt, deployment.StartedAt); err != nil {
		return err
	}
	if err := l.validateReading(ctx, &reading); err != nil {
		return err
	}
	reading.Quality = policy.QualityFor(reading, policy.DefaultLimits())
	return l.recordReading(ctx, reading)
}

func (l *Lab) validateReading(ctx context.Context, reading *model.FluxReading) error {
	if !model.ValidSensorKind(reading.SensorKind) {
		return fmt.Errorf("%w: sensor", model.ErrQuality)
	}
	if err := policy.ValidateFlux(*reading, policy.DefaultLimits()); err != nil {
		return err
	}
	return nil
}
func (l *Lab) recordReading(ctx context.Context, reading model.FluxReading) error {
	if err := l.readings.Save(ctx, &reading); err != nil {
		return err
	}
	l.cache.Update(reading)
	l.metrics.Add("readings_total", 1)
	if reading.Quality == model.QualityGood {
		l.metrics.Add("readings_good", 1)
	}
	return nil
}
func (l *Lab) BatchIngest(ctx context.Context, readings []model.FluxReading) error {
	for _, reading := range readings {
		if err := cancelled(ctx); err != nil {
			return err
		}
		if err := l.IngestReading(ctx, reading); err != nil {
			return err
		}
	}
	return nil
}
func (l *Lab) RecentReadings(ctx context.Context, deploymentID string) ([]model.FluxReading, error) {
	values := l.cache.Sorted(deploymentID)
	if len(values) == 0 {
		return l.readings.ListByDeployment(ctx, deploymentID)
	}
	sort.SliceStable(values, func(i, j int) bool { return values[i].SampledAt.Before(values[j].SampledAt) })
	return values, nil
}
func (l *Lab) ListStationReadings(ctx context.Context, stationID string) ([]model.FluxReading, error) {
	return l.readings.ListByStation(ctx, stationID)
}
