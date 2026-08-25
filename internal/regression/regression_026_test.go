package regression

import (

	"context"
	"errors"
	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/ingest"

)


func TestBug26_RetryAndBatchStopOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()); cancel(); started := time.Now(); err := ingest.Retry(ctx, 3, func() error { return errors.New("temporary") }); if !errors.Is(err, context.Canceled) || time.Since(started) > 30*time.Millisecond { t.Fatalf("retry cancellation err=%v elapsed=%s", err, time.Since(started)) }; results := ingest.RunBatch(ctx, []string{"a", "b"}, func(context.Context, string) error { return nil }); if len(results) != 1 || !errors.Is(results[0].Err, context.Canceled) { t.Fatalf("batch cancellation results=%+v", results) }
}

