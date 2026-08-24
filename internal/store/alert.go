package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type QualityAlertRepo struct{ store *Store }

func NewQualityAlertRepo(s *Store) *QualityAlertRepo { return &QualityAlertRepo{store: s} }
func (r *QualityAlertRepo) Get(ctx context.Context, id string) (*model.QualityAlert, error) {
	return decodeOne[model.QualityAlert](ctx, r.store, "alert", id)
}
func (r *QualityAlertRepo) Save(ctx context.Context, value *model.QualityAlert) error {
	return r.store.SaveContext(ctx, "alert", value.ID, value)
}
func (r *QualityAlertRepo) List(ctx context.Context) ([]model.QualityAlert, error) {
	return decodeMany[model.QualityAlert](ctx, r.store, "alert")
}
func (r *QualityAlertRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "alert", id)
}

func (r *QualityAlertRepo) ListByStation(ctx context.Context, stationID string) ([]model.QualityAlert, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.QualityAlert, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
