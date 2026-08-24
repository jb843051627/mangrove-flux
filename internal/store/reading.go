package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type FluxReadingRepo struct{ store *Store }

func NewFluxReadingRepo(s *Store) *FluxReadingRepo { return &FluxReadingRepo{store: s} }
func (r *FluxReadingRepo) Get(ctx context.Context, id string) (*model.FluxReading, error) {
	return decodeOne[model.FluxReading](ctx, r.store, "reading", id)
}
func (r *FluxReadingRepo) Save(ctx context.Context, value *model.FluxReading) error {
	return r.store.SaveContext(ctx, "reading", value.ID, value)
}
func (r *FluxReadingRepo) List(ctx context.Context) ([]model.FluxReading, error) {
	return decodeMany[model.FluxReading](ctx, r.store, "reading")
}
func (r *FluxReadingRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "reading", id)
}

func (r *FluxReadingRepo) ListByDeployment(ctx context.Context, deploymentID string) ([]model.FluxReading, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.FluxReading, 0, len(items))
	for _, item := range items {
		if item.DeploymentID == deploymentID {
			out = append(out, item)
		}
	}
	return out, nil
}
func (r *FluxReadingRepo) ListByStation(ctx context.Context, stationID string) ([]model.FluxReading, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.FluxReading, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
