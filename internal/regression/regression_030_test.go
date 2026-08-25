package regression

import (

	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/clock"

)


func TestBug30_DayKeyUsesStationTimezone(t *testing.T) { loc, err := time.LoadLocation("Asia/Shanghai"); if err != nil { t.Fatal(err) }; value := time.Date(2026, 4, 1, 23, 30, 0, 0, time.UTC); if got := clock.DayKey(value, loc); got != "2026-04-02" { t.Fatalf("station day = %s", got) } }

