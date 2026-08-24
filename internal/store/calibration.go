package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type CalibrationProfileRepo struct{ store *Store }

func NewCalibrationProfileRepo(s *Store) *CalibrationProfileRepo {
	return &CalibrationProfileRepo{store: s}
}
func (r *CalibrationProfileRepo) Get(ctx context.Context, id string) (*model.CalibrationProfile, error) {
	return decodeOne[model.CalibrationProfile](ctx, r.store, "calibration", id)
}
func (r *CalibrationProfileRepo) Save(ctx context.Context, value *model.CalibrationProfile) error {
	return r.store.SaveContext(ctx, "calibration", value.ID, value)
}
func (r *CalibrationProfileRepo) List(ctx context.Context) ([]model.CalibrationProfile, error) {
	return decodeMany[model.CalibrationProfile](ctx, r.store, "calibration")
}
func (r *CalibrationProfileRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "calibration", id)
}

func (r *CalibrationProfileRepo) ListBySensor(ctx context.Context, sensorKind, stationID string) ([]model.CalibrationProfile, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.CalibrationProfile, 0, len(items))
	for _, item := range items {
		if item.SensorKind == sensorKind && item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
