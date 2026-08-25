package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type WeatherObservationRepo struct{ store *Store }

func NewWeatherObservationRepo(s *Store) *WeatherObservationRepo {
	return &WeatherObservationRepo{store: s}
}
func (r *WeatherObservationRepo) Get(ctx context.Context, id string) (*model.WeatherObservation, error) {
	return decodeOne[model.WeatherObservation](ctx, r.store, "weather", id)
}
func (r *WeatherObservationRepo) Save(ctx context.Context, value *model.WeatherObservation) error {
	return r.store.SaveContext(ctx, "weather", value.ID, value)
}
func (r *WeatherObservationRepo) List(ctx context.Context) ([]model.WeatherObservation, error) {
	return decodeMany[model.WeatherObservation](ctx, r.store, "weather")
}
func (r *WeatherObservationRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "weather", id)
}

func (r *WeatherObservationRepo) ListByStation(ctx context.Context, stationID string) ([]model.WeatherObservation, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.WeatherObservation, 0, len(items))
	for _, item := range items {
		if item.StationID == stationID {
			out = append(out, item)
		}
	}
	return out, nil
}
