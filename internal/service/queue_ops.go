package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/ingest"
)

func (l *Lab) EnqueueQuality(ctx context.Context, readingID string, done chan error) error {
	return l.queue.Submit(ctx, ingest.Job{ID: readingID, Ctx: ctx, Done: done, Run: func(jobCtx context.Context) error { _, err := l.EvaluateReading(jobCtx, readingID); return err }})
}
func (l *Lab) EnqueueSweep(ctx context.Context, stationID string, done chan error) error {
	return l.queue.Submit(ctx, ingest.Job{ID: stationID, Ctx: ctx, Done: done, Run: func(jobCtx context.Context) error { _, err := l.QualitySweep(jobCtx, stationID); return err }})
}
