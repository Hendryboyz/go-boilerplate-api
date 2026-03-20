package params

import (
	"time"
)

func ParseDateString(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func ParseDateTimeString(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date)
}

func ParseFromDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
