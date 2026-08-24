package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/clock"
	"github.com/jb843051627/mangrove-flux/internal/ingest"
	"github.com/jb843051627/mangrove-flux/internal/metrics"
	"github.com/jb843051627/mangrove-flux/internal/store"
)

type Lab struct {
	db           *store.Store
	stations     *store.FieldStationRepo
	plots        *store.PlotRepo
	chambers     *store.ChamberRepo
	deployments  *store.DeploymentRepo
	readings     *store.FluxReadingRepo
	tides        *store.TideObservationRepo
	weather      *store.WeatherObservationRepo
	calibrations *store.CalibrationProfileRepo
	alerts       *store.QualityAlertRepo
	batches      *store.SurveyBatchRepo
	reports      *store.FluxReportRepo
	queue        *ingest.Queue
	metrics      *metrics.Registry
	cache        *ReadingCache
	clock        clock.Clock
}

func NewLab(db *store.Store) *Lab {
	return &Lab{db: db, stations: store.NewFieldStationRepo(db), plots: store.NewPlotRepo(db), chambers: store.NewChamberRepo(db), deployments: store.NewDeploymentRepo(db), readings: store.NewFluxReadingRepo(db), tides: store.NewTideObservationRepo(db), weather: store.NewWeatherObservationRepo(db), calibrations: store.NewCalibrationProfileRepo(db), alerts: store.NewQualityAlertRepo(db), batches: store.NewSurveyBatchRepo(db), reports: store.NewFluxReportRepo(db), queue: ingest.New(32), metrics: metrics.New(), cache: NewReadingCache(), clock: clock.System{}}
}

func (l *Lab) Close() {
	if l.queue != nil {
		l.queue.Close()
	}
}
func (l *Lab) QueueDepth() int                { return l.queue.Pending() }
func (l *Lab) Metrics() map[string]int64      { return l.metrics.Snapshot() }
func (l *Lab) Ping(ctx context.Context) error { return ctx.Err() }
