package regression

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"github.com/jb843051627/mangrove-flux/internal/store"
)

func openBug24(t *testing.T) (*service.Lab, model.Deployment) {
	t.Helper()
	db, err := store.Open(filepath.Join(t.TempDir(), "case.db"))
	if err != nil { t.Fatal(err) }
	app := service.NewLab(db)
	t.Cleanup(func() { app.Close(); db.Close() })
	now := time.Now().UTC()
	station := model.FieldStation{ID: "station-24", Name: "东湾红树林", Region: "delta", Timezone: "Asia/Shanghai", TideDatum: "local", Active: true, CreatedAt: now}
	if err := app.CreateStation(context.Background(), station); err != nil { t.Fatal(err) }
	plot := model.Plot{ID: "plot-24", StationID: station.ID, Name: "潮沟边缘", Habitat: "fringe", AreaM2: 18, TargetDepthM: 0.4}
	if err := app.CreatePlot(context.Background(), plot); err != nil { t.Fatal(err) }
	chamber := model.Chamber{ID: "chamber-24", StationID: station.ID, Serial: "CH-24", VolumeL: 42, SensorKind: "co2", InstalledAt: now}
	if err := app.RegisterChamber(context.Background(), chamber); err != nil { t.Fatal(err) }
	batch := model.SurveyBatch{ID: "batch-24", StationID: station.ID, Name: "四月潮窗", StartedAt: now, ExpectedDeployments: 1}
	if err := app.OpenBatch(context.Background(), batch); err != nil { t.Fatal(err) }
	deployment := model.Deployment{ID: "deployment-24", BatchID: batch.ID, PlotID: plot.ID, ChamberID: chamber.ID, Operator: "field-team"}
	if err := app.PlanDeployment(context.Background(), deployment); err != nil { t.Fatal(err) }
	if err := app.StartDeployment(context.Background(), deployment.ID); err != nil { t.Fatal(err) }
	return app, deployment
}

func requireNotFound24(t *testing.T, err error) {
	t.Helper()
	if !errors.Is(err, model.ErrNotFound) { t.Fatalf("expected not-found, got %v", err) }
}


func TestBug24_ClearedAlertCannotBeReviewed(t *testing.T) {
	app, _ := openBug24(t); alert := model.QualityAlert{ID: "alert-24", StationID: "station-24", Code: "quality", Severity: "warning", State: model.AlertCleared, Message: "已清理", CreatedAt: time.Now().UTC()}; if err := app.RecordAlert(context.Background(), alert); err != nil { t.Fatal(err) }; if err := app.ReviewAlert(context.Background(), alert.ID); !errors.Is(err, model.ErrConflict) { t.Fatalf("cleared alert review returned %v", err) }
}

