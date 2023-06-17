package main

import (
	"math"
	"time"
)

// A CivilTime represents an instant in time designated by civilian authorities
// with at least millisecond precision.
type CivilTime struct {
	local time.Time
}

// NewCivilTime creates new CivilTime corresponding to the given time in
// proleptic Gregorian calendar.
func NewCivilTime(t time.Time) CivilTime {
	return CivilTime{local: t}
}

// Date returns the year, month, and decimal day in which ct occurs (in used
// time zone).
func (ct CivilTime) Date() (year, month, day float64) {
	y, m, _ := ct.local.Date()
	year = float64(y)
	month = float64(m)
	day = decimalDay(ct.local)
	return year, month, day
}

// Equal reports whether ct and ct2 represent the same time instant.
func (ct CivilTime) Equal(ct2 CivilTime) bool {
	return ct.local.Equal(ct2.local)
}

// JulianDay returns number of days counted from -4713-11-24T12:00Z (noon of
// 24 November -4714 BC) in proleptic Gregorian calendar (in Julian calendar it
// is -4712-01-01T12:00Z or noon of January 1, 4713 BC).
//
// Method mostly implements algorithm from J.L. Lawrence "Celestial Calculations"
// (2019) which is improved version of algorithm from J. Meeus "Astronomical
// Algorithms" (1998). Differences in current implementation:
// - in Go constant numbers are not rounded, so equation can use 30.6 instead of
//   30.6001
// - removed correction for Julian calendar.
func (ct CivilTime) JulianDay() float64 {
	y, m, d := ct.ToUTC().Date()
	if m <= 2 {
		y -= 1
		m += 12
	}
	T := 0.0
	if y < 0 {
		T = 0.75
	}
	A := math.Trunc(y / 100)
	B := 2 - A + math.Trunc(A/4)
	JD := B + math.Trunc(365.25*y-T) + math.Trunc(30.6*(m+1)) + d + 1720994.5
	return JD
}

// ToUTC returns copy of ct representing the same time instant, but with the
// copy's location information set to UTC for display purposes.
func (ct CivilTime) ToUTC() CivilTime {
	return CivilTime{local: ct.local.UTC()}
}
