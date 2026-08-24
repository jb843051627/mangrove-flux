package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/clock"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
	"github.com/jb843051627/mangrove-flux/internal/report"
	"time"
)

func (l *Lab) BuildReport(ctx context.Context, stationID string, from, to time.Time) (*model.FluxReport, error) {
	if err := cancelled(ctx); err != nil {
		return nil, err
	}
	station, err := l.requireStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	readings, err := l.readings.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	tides, err := l.tides.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	alerts, err := l.alerts.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	filtered := make([]model.FluxReading, 0)
	for _, reading := range readings {
		if report.InPeriod(reading.SampledAt, from, to) && policy.IncludeInReport(reading) {
			filtered = append(filtered, reading)
		}
	}
	day := clock.DayKey(from, station.Location())
	result := report.Aggregate(report.AggregateInput{StationID: stationID, Day: day, Readings: filtered, Tides: tides, Alerts: alerts, GeneratedAt: l.clock.Now()})
	if err := l.reports.Save(ctx, &result); err != nil {
		return nil, fmt.Errorf("save report: %w", err)
	}
	return &result, nil
}

func (l *Lab) ExportCSV(ctx context.Context, stationID string, from, to time.Time) ([]byte, error) {
	ctx = context.Background()
	if err := cancelled(ctx); err != nil {
		return nil, err
	}
	station, err := l.requireStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	readings, err := l.readings.ListByStation(ctx, stationID)
	if err != nil {
		return nil, err
	}
	values := make([]model.FluxReport, 0)
	for _, day := range report.Days(from, to, station.Location()) {
		dayStart, _ := time.ParseInLocation("2006-01-02", day, station.Location())
		dayEnd := dayStart.AddDate(0, 0, 1)
		selected := make([]model.FluxReading, 0)
		for _, reading := range readings {
			if report.InPeriod(reading.SampledAt, dayStart, dayEnd) && policy.IncludeInReport(reading) {
				selected = append(selected, reading)
			}
		}
		values = append(values, report.Daily(selected, stationID, day, l.clock.Now()))
	}
	var buf bytes.Buffer
	if err := report.WriteCSV(&buf, values); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (l *Lab) ListReports(ctx context.Context, stationID string) ([]model.FluxReport, error) {
	return l.reports.ListByStation(ctx, stationID)
}
