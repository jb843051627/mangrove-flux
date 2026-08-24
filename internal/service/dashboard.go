package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

type Dashboard struct {
	Station           model.FieldStation `json:"station"`
	Plots             int                `json:"plots"`
	Chambers          int                `json:"chambers"`
	ActiveDeployments int                `json:"active_deployments"`
	Readings          int64              `json:"readings"`
	GoodReadings      int64              `json:"good_readings"`
	Alerts            int                `json:"alerts"`
	QueueDepth        int                `json:"queue_depth"`
	Metrics           map[string]int64   `json:"metrics"`
}

func (l *Lab) DashboardSnapshot(ctx context.Context, stationID string) (*Dashboard, error) {
	station, err := l.requireStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	plots, err := l.plots.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	chambers, err := l.chambers.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	alerts, err := l.alerts.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	deployments, err := l.deployments.List(ctx)
	if err != nil {
		return nil, err
	}
	active := 0
	for _, deployment := range deployments {
		if deployment.IsActive() {
			active++
		}
	}
	values := l.metrics.Snapshot()
	return &Dashboard{Station: *station, Plots: len(plots), Chambers: len(chambers), ActiveDeployments: active, Readings: values["readings_total"], GoodReadings: values["readings_good"], Alerts: len(alerts), QueueDepth: l.queue.Pending(), Metrics: values}, nil
}
