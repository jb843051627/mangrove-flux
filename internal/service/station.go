package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"time"
)

func (l *Lab) CreateStation(ctx context.Context, station model.FieldStation) error {
	if err := station.Validate(); err != nil {
		return err
	}
	if station.CreatedAt.IsZero() {
		station.CreatedAt = l.clock.Now()
	}
	return l.stations.Save(ctx, &station)
}
func (l *Lab) GetStation(ctx context.Context, id string) (*model.FieldStation, error) {
	return l.stations.Get(ctx, id)
}
func (l *Lab) ListStations(ctx context.Context) ([]model.FieldStation, error) {
	return l.stations.List(ctx)
}
func (l *Lab) DisableStation(ctx context.Context, id string) error {
	station, err := l.stations.Get(ctx, id)
	if err != nil {
		return err
	}
	station.Active = false
	return l.stations.Save(ctx, station)
}
func (l *Lab) StationAge(station model.FieldStation, now time.Time) time.Duration {
	return now.Sub(station.CreatedAt)
}
