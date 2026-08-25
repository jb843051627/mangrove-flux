package regression

import (

	"context"
	"errors"
	"path/filepath"
	"testing"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"github.com/jb843051627/mangrove-flux/internal/store"

)


func TestBug22_EmptyTideSeriesReturnsNotFound(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "case.db")); if err != nil { t.Fatal(err) }; app := service.NewLab(db); defer app.Close(); defer db.Close(); _, err = app.TideHeight(context.Background(), "station-22"); if !errors.Is(err, model.ErrNotFound) { t.Fatalf("empty tide series returned %v", err) }
}

