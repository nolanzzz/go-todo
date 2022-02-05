package utils

import (
	"time"
)

func TimeDiffMinutes(start, end time.Time) (minutes int) {
	minutes = int(end.Sub(start).Minutes())
	return minutes
}

func TimeUntilDayEnd() time.Duration {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	tomorrowStart := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.Local)
	diff := tomorrowStart.Sub(now) * time.Second
	return diff
}
