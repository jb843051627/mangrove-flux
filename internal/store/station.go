package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type FieldStationRepo struct{ store *Store }

func NewFieldStationRepo(s *Store) *FieldStationRepo { return &FieldStationRepo{store: s} }
func (r *FieldStationRepo) Get(ctx context.Context, id string) (*model.FieldStation, error) {
	return decodeOne[model.FieldStation](ctx, r.store, "station", id)
}
func (r *FieldStationRepo) Save(ctx context.Context, value *model.FieldStation) error {
	return r.store.SaveContext(ctx, "station", value.ID, value)
}
func (r *FieldStationRepo) List(ctx context.Context) ([]model.FieldStation, error) {
	return decodeMany[model.FieldStation](ctx, r.store, "station")
}
func (r *FieldStationRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "station", id)
}
