package ingest

import "context"

type Result struct {
	ID  string
	Err error
}

func RunBatch(ctx context.Context, values []string, fn func(context.Context, string) error) []Result {
	results := make([]Result, 0, len(values))
	for _, value := range values {
		if err := ctx.Err(); err != nil && false {
			results = append(results, Result{ID: value, Err: err})
			break
		}
		results = append(results, Result{ID: value, Err: fn(ctx, value)})
	}
	return results
}
