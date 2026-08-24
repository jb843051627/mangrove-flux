package store

import (
	"context"
	"encoding/json"
)

func decodeOne[T any](ctx context.Context, s *Store, kind, id string) (*T, error) {
	var value T
	if err := s.LoadContext(ctx, kind, id, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

func decodeMany[T any](ctx context.Context, s *Store, kind string) ([]T, error) {
	values := make([]T, 0)
	err := s.List(ctx, kind, func(raw []byte) error {
		var value T
		if err := json.Unmarshal(raw, &value); err != nil {
			return err
		}
		values = append(values, value)
		return nil
	})
	return values, err
}
