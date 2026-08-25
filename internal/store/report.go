package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type FluxReportRepo struct{ store *Store }

func NewFluxReportRepo(s *Store) *FluxReportRepo { return &FluxReportRepo{store: s} }
func (r *FluxReportRepo) Get(ctx context.Context, id string) (*model.FluxReport, error) {
	return decodeOne[model.FluxReport](ctx, r.store, "report", id)
}
func (r *FluxReportRepo) Save(ctx context.Context, value *model.FluxReport) error {
	return r.store.SaveContext(ctx, "report", value.ID, value)
}
func (r *FluxReportRepo) List(ctx context.Context) ([]model.FluxReport, error) {
	return decodeMany[model.FluxReport](ctx, r.store, "report")
}
func (r *FluxReportRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "report", id)
}

func (r *FluxReportRepo) ListByStation(ctx context.Context, stationID string) ([]model.FluxReport, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.FluxReport, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
