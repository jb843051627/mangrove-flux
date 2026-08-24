package store

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type DeploymentRepo struct{ store *Store }

func NewDeploymentRepo(s *Store) *DeploymentRepo { return &DeploymentRepo{store: s} }
func (r *DeploymentRepo) Get(ctx context.Context, id string) (*model.Deployment, error) {
	return decodeOne[model.Deployment](ctx, r.store, "deployment", id)
}
func (r *DeploymentRepo) Save(ctx context.Context, value *model.Deployment) error {
	return r.store.SaveContext(ctx, "deployment", value.ID, value)
}
func (r *DeploymentRepo) List(ctx context.Context) ([]model.Deployment, error) {
	return decodeMany[model.Deployment](ctx, r.store, "deployment")
}
func (r *DeploymentRepo) Delete(ctx context.Context, id string) error {
	return r.store.DeleteContext(ctx, "deployment", id)
}

func (r *DeploymentRepo) ListByBatch(ctx context.Context, batchID string) ([]model.Deployment, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.Deployment, 0, len(items))
	for _, item := range items {
		if item.BatchID == batchID {
			out = append(out, item)
		}
	}
	return out, nil
}
func (r *DeploymentRepo) ListByPlot(ctx context.Context, plotID string) ([]model.Deployment, error) {
	items, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]model.Deployment, 0, len(items))
	for _, item := range items {
		if item.PlotID == plotID {
			out = append(out, item)
		}
	}
	return out, nil
}
