package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type TideObservationRepo struct{ store *Store }

func NewTideObservationRepo(s *Store) *TideObservationRepo { return &TideObservationRepo{store: s} }
func (r *TideObservationRepo) Get(ctx context.Context, id string) (*model.TideObservation, error) {
	return decodeOne[model.TideObservation](ctx, r.store, "tide", id)
}
func (r *TideObservationRepo) Save(ctx context.Context, value *model.TideObservation) error {
	return r.store.SaveContext(ctx, "tide", value.ID, value)
}
func (r *TideObservationRepo) List(ctx context.Context) ([]model.TideObservation, error) {
	return decodeMany[model.TideObservation](ctx, r.store, "tide")
}
func (r *TideObservationRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "tide", id)
}

func (r *TideObservationRepo) ListByStation(ctx context.Context, stationID string) ([]model.TideObservation, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.TideObservation, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
