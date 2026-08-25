package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type ChamberRepo struct{ store *Store }

func NewChamberRepo(s *Store) *ChamberRepo { return &ChamberRepo{store: s} }
func (r *ChamberRepo) Get(ctx context.Context, id string) (*model.Chamber, error) {
	return decodeOne[model.Chamber](ctx, r.store, "chamber", id)
}
func (r *ChamberRepo) Save(ctx context.Context, value *model.Chamber) error {
	return r.store.SaveContext(ctx, "chamber", value.ID, value)
}
func (r *ChamberRepo) List(ctx context.Context) ([]model.Chamber, error) {
	return decodeMany[model.Chamber](ctx, r.store, "chamber")
}
func (r *ChamberRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "chamber", id)
}

func (r *ChamberRepo) ListByStation(ctx context.Context, stationID string) ([]model.Chamber, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.Chamber, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
