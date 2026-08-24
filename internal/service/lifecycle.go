package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"time"
)

type HealthStatus struct {
	Stations  int       `json:"stations"`
	Plots     int       `json:"plots"`
	Readings  int64     `json:"readings"`
	Queue     int       `json:"queue"`
	CheckedAt time.Time `json:"checked_at"`
}

func (l *Lab) Health(ctx context.Context) (*HealthStatus, error) {
	if err := cancelled(ctx); err != nil {
		return nil, err
	}
	stations, err := l.stations.List(ctx)
	if err != nil {
		return nil, err
	}
	plots, err := l.plots.List(ctx)
	if err != nil {
		return nil, err
	}
	return &HealthStatus{Stations: len(stations), Plots: len(plots), Readings: l.metrics.Get("readings_total"), Queue: l.queue.Pending(), CheckedAt: l.clock.Now()}, nil
}
func (l *Lab) CloseIdleDeployments(ctx context.Context, before time.Time) (int, error) {
	values, err := l.deployments.List(ctx)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, value := range values {
		if value.State == model.DeploymentRunning && !value.StartedAt.After(before) {
			if err := l.CloseDeployment(ctx, value.ID); err == nil {
				count++
			}
		}
	}
	return count, nil
}
func (l *Lab) PruneReadingCache(before time.Time) { l.cache.PruneBefore(before) }
