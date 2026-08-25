package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
)

func (l *Lab) CreatePlot(ctx context.Context, plot model.Plot) error {
	if err := plot.Validate(); err != nil {
		return err
	}
	station, err := l.requireStation(ctx, plot.StationID)
	if err != nil {
		return err
	}
	if !policy.StationCanReceive(*station) {
		return model.ErrClosed
	}
	plot.Active = true
	return l.plots.Save(ctx, &plot)
}
func (l *Lab) GetPlot(ctx context.Context, id string) (*model.Plot, error) {
	return l.plots.Get(ctx, id)
}
func (l *Lab) ListPlots(ctx context.Context, stationID string) ([]model.Plot, error) {
	if _, err := l.requireStation(ctx, stationID); err != nil {
		return nil, err
	}
	return l.plots.ListByStation(ctx, stationID)
}
func (l *Lab) DeactivatePlot(ctx context.Context, id string) error {
	plot, err := l.plots.Get(ctx, id)
	if err != nil {
		return err
	}
	plot.Active = false
	return l.plots.Save(ctx, plot)
}
