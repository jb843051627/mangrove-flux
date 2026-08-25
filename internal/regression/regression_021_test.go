package regression

import (

	"context"
	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/ingest"

)


func TestBug21_SubmitAfterQueueCloseReturnsImmediately(t *testing.T) {
	q := ingest.New(1); q.Close(); ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond); defer cancel(); started := time.Now(); err := q.Submit(ctx, ingest.Job{Run: func(context.Context) error { return nil }}); if err == nil || time.Since(started) > 40*time.Millisecond { t.Fatalf("submit after close returned err=%v after %s", err, time.Since(started)) }
}

