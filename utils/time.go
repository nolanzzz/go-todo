package utils

import (
	"time"
)

func TimeDiffSeconds(start, end time.Time) (seconds int) {
	seconds = int(end.Unix() - start.Unix())
	return seconds
}
