package utils

import "time"

func TimeDiffSeconds(start, end time.Time) (seconds int) {
	seconds = end.Second() - start.Second()
	return seconds
}
