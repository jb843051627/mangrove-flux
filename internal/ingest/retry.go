package ingest

import (
	"context"
	"time"
)

func Retry(ctx context.Context, attempts int, fn func() error) error {
	if attempts < 1 {
		attempts = 1
	}
	var last error
	for i := 0; i < attempts; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		last = fn()
		if last == nil {
			return nil
		}
		timer := time.NewTimer(time.Duration(i+1) * 5 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
	return last
}
