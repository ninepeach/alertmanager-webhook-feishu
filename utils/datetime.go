package utils

import (
	"time"
)

// PrevOfMonth
func PrevOfMonth(ts time.Time) time.Time {
	return BeginningOfMonth(ts).AddDate(0, -1, 0)
}

// BeginningOfMonth beginning of month
func BeginningOfMonth(ts time.Time) time.Time {
	y, m, _ := ts.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, ts.Location())
}

// EndOfMonth end of month
func EndOfMonth(ts time.Time) time.Time {
	y, m, _ := ts.Date()
	return time.Date(y, m+1, 1, 0, 0, 0, 0, ts.Location())
}
