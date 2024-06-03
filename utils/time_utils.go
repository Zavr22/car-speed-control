package utils

import "time"

func IsWithinAccessHours(start, end time.Time) bool {
	now := time.Now()

	startToday := time.Date(now.Year(), now.Month(), now.Day(), start.Hour(), start.Minute(), 0, 0, now.Location())

	endToday := time.Date(now.Year(), now.Month(), now.Day(), end.Hour(), end.Minute(), 0, 0, now.Location())

	return now.After(startToday) && now.Before(endToday)
}
