package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
	"github.com/jb843051627/mangrove-flux/internal/store"
)

func (l *Lab) OpenBatch(ctx context.Context, batch model.SurveyBatch) error {
	if err := batch.Validate(); err != nil {
		return err
	}
	if _, err := l.requireStation(ctx, batch.StationID); err != nil {
		return err
	}
	batch.State = model.BatchOpen
	return l.batches.Save(ctx, &batch)
}
func (l *Lab) GetBatch(ctx context.Context, id string) (*model.SurveyBatch, error) {
	return l.batches.Get(ctx, id)
}
func (l *Lab) ListBatches(ctx context.Context) ([]model.SurveyBatch, error) {
	return l.batches.List(ctx)
}

func (l *Lab) OpenBatchWithDeployments(ctx context.Context, batch model.SurveyBatch, deployments []model.Deployment) error {
	if err := batch.Validate(); err != nil {
		return err
	}
	if _, err := l.requireStation(ctx, batch.StationID); err != nil {
		return err
	}
	batch.State = model.BatchOpen
	return l.db.WithTx(ctx, func(tx *sql.Tx) error {
		if err := store.TxSave(tx, "batch", batch.ID, &batch); err != nil {
			return err
		}
		for _, deployment := range deployments {
			if err := deployment.Validate(); err != nil {
				return fmt.Errorf("deployment %s: %w", deployment.ID, err)
			}
			if err := store.TxSave(tx, "deployment", deployment.ID, &deployment); err != nil {
				return err
			}
		}
		return nil
	})
}

func (l *Lab) closeBatch(ctx context.Context, batch *model.SurveyBatch, deployments []model.Deployment) error {
	if !policy.BatchCanClose(*batch, deployments) {
		return model.ErrConflict
	}
	if !model.CanBatchTransition(batch.State, model.BatchClosed) {
		return model.ErrConflict
	}
	now := l.clock.Now()
	batch.State = model.BatchClosed
	batch.ClosedAt = &now
	return l.batches.Save(ctx, batch)
}

func (l *Lab) CloseBatch(ctx context.Context, id string) error {
	batch, err := l.batches.Get(ctx, id)
	if err != nil {
		return err
	}
	deployments, err := l.deployments.ListByBatch(ctx, id)
	if err != nil {
		return err
	}
	return l.closeBatch(ctx, batch, deployments)
}
func (l *Lab) AbortBatch(ctx context.Context, id string) error {
	batch, err := l.batches.Get(ctx, id)
	if err != nil {
		return err
	}
	if !model.CanBatchTransition(batch.State, model.BatchAborted) {
		return model.ErrConflict
	}
	batch.State = model.BatchAborted
	return l.batches.Save(ctx, batch)
}
func (l *Lab) CompletionRatio(ctx context.Context, id string) (float64, error) {
	batch, err := l.batches.Get(ctx, id)
	if err != nil {
		return 0, err
	}
	return policy.CompletionRatio(*batch), nil
}

func (l *Lab) UpdateBatchRevision(ctx context.Context, batch model.SurveyBatch, expected int) error {
	return l.batches.Save(ctx, &batch)
}

var _ = store.TxSave
