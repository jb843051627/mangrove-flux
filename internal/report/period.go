package report

import (
	"github.com/jb843051627/mangrove-flux/internal/clock"
	"time"
)

func Days(from, to time.Time, loc *time.Location) []string {
	if to.Before(from) {
		return nil
	}
	start := clock.DayStart(from, loc)
	end := clock.DayStart(to, loc)
	out := make([]string, 0)
	for current := start; !current.After(end); current = current.AddDate(0, 0, 1) {
		out = append(out, clock.DayKey(current, loc))
	}
	return out
}

func InPeriod(value, from, to time.Time) bool { return !value.Before(from) && value.Before(to) }
