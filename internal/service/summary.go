package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"sort"
)

func (l *Lab) DeploymentSummary(ctx context.Context, batchID string) (map[string]int, error) {
	values, err := l.deployments.ListByBatch(ctx, batchID)
	if err != nil {
		return nil, err
	}
	result := map[string]int{"planned": 0, "running": 0, "closed": 0, "void": 0}
	for _, value := range values {
		result[string(value.State)]++
	}
	return result, nil
}
func (l *Lab) SortedStations(ctx context.Context) ([]model.FieldStation, error) {
	values, err := l.stations.List(ctx)
	if err != nil {
		return nil, err
	}
	sort.SliceStable(values, func(i, j int) bool { return values[i].Name < values[j].Name })
	return values, nil
}
