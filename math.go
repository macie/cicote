package main

import (
	"math"
	"time"
)

// decimalDay returns given day with time in decimal format.
func decimalDay(t time.Time) float64 {
	hour, min, sec := t.Clock()
	return float64(t.Day()) +
		float64(hour)/24 + float64(min)/1440 + float64(sec)/8.64e4 +
		float64(t.Nanosecond())/8.64e13
}

// decimalHour returns given time (without date) in decimal format.
func decimalHour(t time.Time) float64 {
	hour, min, sec := t.Clock()
	return float64(hour) + float64(min)/60 + float64(sec)/3.6e3 +
		float64(t.Nanosecond())/3.6e12
}

// round returns number with given precision.
//
// The 64-bit floating-point numbers have precision of 15 or 16 significant
// digits so calling round with precision >15 for some numbers may return
// invalid result.
func round(f float64, precision uint8) float64 {
	p := math.Pow10(int(precision))
	return math.Round(f*p) / p
}
