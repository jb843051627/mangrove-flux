package regression

import (

	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/report"

)


func TestBug25_EmptyDailyReportHasFiniteTotals(t *testing.T) { value := report.Daily(nil, "station-25", "2026-04-01", time.Now().UTC()); if !model.Finite(value.NetCarbon) || !model.Finite(value.MeanCO2) { t.Fatalf("empty report contains non-finite values: %+v", value) } }

