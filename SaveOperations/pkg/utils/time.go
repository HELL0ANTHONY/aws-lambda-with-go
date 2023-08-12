package utils

import "time"

func Time[T int | string](t func(time.Time) T) T {
	now := time.Now()
	location, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	timestamp := now.In(location)
	return t(timestamp)
}
