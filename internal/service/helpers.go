package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

func cancelled(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("%w: %v", model.ErrCancelled, err)
	}
	return nil
}
func wrap(kind string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", kind, err)
}
func isMissing(err error) bool { return errors.Is(err, model.ErrNotFound) }

func (l *Lab) requireStation(ctx context.Context, id string) (*model.FieldStation, error) {
	station, err := l.stations.Get(ctx, id)
	if err != nil {
		return nil, wrap("station", err)
	}
	if !station.Active {
		return nil, model.ErrClosed
	}
	return station, nil
}
func (l *Lab) requirePlot(ctx context.Context, id string) (*model.Plot, error) {
	plot, err := l.plots.Get(ctx, id)
	if err != nil {
		return nil, wrap("plot", err)
	}
	return plot, nil
}
func (l *Lab) requireDeployment(ctx context.Context, id string) (*model.Deployment, error) {
	deployment, err := l.deployments.Get(ctx, id)
	if err != nil {
		return nil, wrap("deployment", err)
	}
	return deployment, nil
}
