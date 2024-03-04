package utils

import (
	"time"
)

const (
	ISTSuffix    = " IST"
	DateTimeZone = "2006-01-02 15:04:05 MST"
)

// ParseDateTime : Allowed format examples : "2023-12-21 14:45:40"
func ParseDateTime(timeStr string) (time.Time, error) {
	date, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

// ParseDateTimeZone : Allowed format examples : "2023-12-21 14:45:40 IST"
func ParseDateTimeZone(timeStr string) (time.Time, error) {
	date, err := time.Parse(DateTimeZone, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func ConvertDateToUTC(date time.Time) time.Time {
	return date.UTC()
}

func ConvertDateToIST(date time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return date.In(loc)
}
