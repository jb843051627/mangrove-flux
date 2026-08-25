package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type SurveyBatchRepo struct{ store *Store }

func NewSurveyBatchRepo(s *Store) *SurveyBatchRepo { return &SurveyBatchRepo{store: s} }
func (r *SurveyBatchRepo) Get(ctx context.Context, id string) (*model.SurveyBatch, error) {
	return decodeOne[model.SurveyBatch](ctx, r.store, "batch", id)
}
func (r *SurveyBatchRepo) Save(ctx context.Context, value *model.SurveyBatch) error {
	return r.store.SaveContext(ctx, "batch", value.ID, value)
}
func (r *SurveyBatchRepo) List(ctx context.Context) ([]model.SurveyBatch, error) {
	return decodeMany[model.SurveyBatch](ctx, r.store, "batch")
}
func (r *SurveyBatchRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "batch", id)
}

func (r *SurveyBatchRepo) UpdateRevision(ctx context.Context, value *model.SurveyBatch, expected int) error {
	current, err := r.Get(ctx, value.ID)
	if err != nil {
		return err
	}
	if current.Revision != expected {
		return model.ErrConflict
	}
	value.Revision = expected + 1
	return r.Save(ctx, value)
}
