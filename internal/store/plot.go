package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type PlotRepo struct{ store *Store }

func NewPlotRepo(s *Store) *PlotRepo { return &PlotRepo{store: s} }
func (r *PlotRepo) Get(ctx context.Context, id string) (*model.Plot, error) {
	return decodeOne[model.Plot](ctx, r.store, "plot", id)
}
func (r *PlotRepo) Save(ctx context.Context, value *model.Plot) error {
	return r.store.SaveContext(ctx, "plot", value.ID, value)
}
func (r *PlotRepo) List(ctx context.Context) ([]model.Plot, error) {
	return decodeMany[model.Plot](ctx, r.store, "plot")
}
func (r *PlotRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "plot", id)
}

func (r *PlotRepo) ListByStation(ctx context.Context, stationID string) ([]model.Plot, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.Plot, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
