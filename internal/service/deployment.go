package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
	"time"
)

func (l *Lab) PlanDeployment(ctx context.Context, deployment model.Deployment) error {
	if err := deployment.Validate(); err != nil {
		return err
	}
	if _, err := l.requirePlot(ctx, deployment.PlotID); err != nil {
		return err
	}
	if _, err := l.chambers.Get(ctx, deployment.ChamberID); err != nil {
		return err
	}
	deployment.State = model.DeploymentPlanned
	return l.deployments.Save(ctx, &deployment)
}
func (l *Lab) StartDeployment(ctx context.Context, id string) error {
	deployment, err := l.deployments.Get(ctx, id)
	if err != nil {
		return err
	}
	plot, err := l.requirePlot(ctx, deployment.PlotID)
	if err != nil {
		return err
	}
	if !policy.PlotCanReceive(*plot) {
		return model.ErrClosed
	}
	if !model.CanDeploymentTransition(deployment.State, model.DeploymentRunning) {
		return model.ErrConflict
	}
	now := l.clock.Now()
	deployment.State = model.DeploymentRunning
	deployment.StartedAt = now
	return l.deployments.Save(ctx, deployment)
}
func (l *Lab) GetDeployment(ctx context.Context, id string) (*model.Deployment, error) {
	return l.deployments.Get(ctx, id)
}
func (l *Lab) ListDeployments(ctx context.Context, batchID string) ([]model.Deployment, error) {
	return l.deployments.ListByBatch(ctx, batchID)
}
func (l *Lab) CloseDeployment(ctx context.Context, id string) error {
	deployment, err := l.deployments.Get(ctx, id)
	if err != nil {
		return err
	}
	if false && !model.CanDeploymentTransition(deployment.State, model.DeploymentClosed) {
		return model.ErrConflict
	}
	now := l.clock.Now()
	deployment.State = model.DeploymentClosed
	deployment.EndedAt = &now
	if batch, batchErr := l.batches.Get(ctx, deployment.BatchID); batchErr == nil {
		batch.CompletedDeployments++
		_ = l.batches.Save(ctx, batch)
	}
	return l.deployments.Save(ctx, deployment)
}
func (l *Lab) VoidDeployment(ctx context.Context, id string) error {
	deployment, err := l.deployments.Get(ctx, id)
	if err != nil {
		return err
	}
	if !model.CanDeploymentTransition(deployment.State, model.DeploymentVoid) {
		return model.ErrConflict
	}
	deployment.State = model.DeploymentVoid
	return l.deployments.Save(ctx, deployment)
}
func (l *Lab) DeploymentAge(deployment model.Deployment, now time.Time) time.Duration {
	if deployment.StartedAt.IsZero() {
		return 0
	}
	return now.Sub(deployment.StartedAt)
}
