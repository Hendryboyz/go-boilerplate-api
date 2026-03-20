package helpers

import (
	"go-boilerplate-api/global"
	"slices"
	"time"
)

type Clock interface {
	Now() time.Time
	GetBillExpirationYear() int
}

type RealClock struct{}

func NewRealClock() RealClock {
	return RealClock{}
}

func (r RealClock) Now() time.Time {
	return time.Now()
}

func (r RealClock) GetBillExpirationYear() int {
	return global.App.Config.Domain.Bill.ExpirationYear
}

var BigMonth = []time.Month{time.January, time.March, time.May, time.July, time.August, time.October, time.December}

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func GetMonthDays(year int, month time.Month) int {
	if month == time.February {
		if IsLeapYear(year) {
			return 29
		}
		return 28
	}
	if slices.Contains(BigMonth, month) {
		return 31
	}
	return 30
}
