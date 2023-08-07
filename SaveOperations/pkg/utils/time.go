package utils

import "time"

type TimeOperation func(time.Time) any

func Time(t TimeOperation) (any, error) {
	now := time.Now()
	location, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return nil, err
	}
	timestamp := now.In(location)
	return t(timestamp), nil
}
