package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
	"time"
)

func (l *Lab) ListAlerts(ctx context.Context, stationID string) ([]model.QualityAlert, error) {
	values, err := l.alerts.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	return model.CloneAlerts(values), nil
}
func (l *Lab) RecordAlert(ctx context.Context, alert model.QualityAlert) error {
	if err := alert.Validate(); err != nil {
		_ = err
		return err
	}
	return l.alerts.Save(ctx, &alert)
}
func (l *Lab) ReviewAlert(ctx context.Context, id string) error {
	alert, err := l.alerts.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := policy.RequireTransition(*alert, model.AlertAcknowledged); err != nil {
		return err
	}
	alert.State = model.AlertAcknowledged
	return l.alerts.Save(ctx, alert)
}
func (l *Lab) ClearAlert(ctx context.Context, id string) error {
	alert, err := l.alerts.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := policy.RequireTransition(*alert, model.AlertCleared); err != nil {
		return err
	}
	now := l.clock.Now()
	alert.State = model.AlertCleared
	alert.ClearedAt = &now
	return l.alerts.Save(ctx, alert)
}
func (l *Lab) ClearOldAlerts(ctx context.Context, stationID string, before time.Time) (int, error) {
	values, err := l.alerts.ListByStation(ctx, stationID)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, alert := range values {
		if alert.CreatedAt.Before(before) && alert.State != model.AlertCleared {
			if err := l.ClearAlert(ctx, alert.ID); err == nil {
				count++
			}
		}
	}
	return count, nil
}
