package clock

import "time"

func DayStart(value time.Time, loc *time.Location) time.Time {
	local := value.In(loc)
	return time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, loc)
}

func DayKey(value time.Time, loc *time.Location) string { return value.In(loc).Format("2006-01-02") }

func InSameDay(a, b time.Time, loc *time.Location) bool { return DayKey(a, loc) == DayKey(b, loc) }
