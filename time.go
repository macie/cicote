package main

import (
	"fmt"
	"math"
	"time"
)

// A CivilTime represents an instant in time designated by civilian authorities.
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

// String returns the time formatted as ISO 8601.
func (ct CivilTime) String() string {
	return ct.local.Format(time.RFC3339)
}

// ToLCT returns copy of ct representing the same time instant, but with the
// copy's location changed by offset for display purposes.
func (ct CivilTime) ToLCT(offset time.Duration) CivilTime {
	zone := time.FixedZone("local", int(math.Round(offset.Seconds())))
	return CivilTime{
		local: ct.local.In(zone),
	}
}

// ToUTC returns copy of ct representing the same time instant, but with the
// copy's location information set to UTC for display purposes.
func (ct CivilTime) ToUTC() CivilTime {
	return CivilTime{local: ct.local.UTC()}
}

// ToGST returns SiderealTime for mean equinox at 0 longitude (Greenwich mean
// sidereal time) representing the same time instant with accuracy ± 1 second.
//
// Using UT1 time instead of UTC may theoretically improve accuracy by multiple
// orders of magnitude.
//
// Method implements algorithm from S. Urban and P.K. Seidelmann "Explanatory
// Supplement to the Astronomical Almanac" (2013) with precession correction
// for dates from range 1000-3000 year.
func (ct CivilTime) ToGST() SiderealTime {
	const dayInSec = 86400
	utc := ct.ToUTC()
	year, month, day := utc.Date()
	date := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

	t_U := utc.JulianDay() - 2451545                            // Julian date from epoch J2000.0
	t := t_U / 36525                                            // Julian centuries from epoch J2000.0
	precession := 0.00096707 + t*(307.47710227+t*(0.092772113)) // -0.000000029 * t^3 + 0.000001997 * t^4
	GST := math.FMA(dayInSec, 0.7790572732640+1.00273781191135448*t_U, precession)
	GST = math.Mod(GST, dayInSec)
	if GST < 0 {
		GST += dayInSec
	}
	return SiderealTime{longitude: 0, date: date, timeInSec: GST}
}

// ToLST returns SiderealTime for mean equinox at given longitude (local mean
// sidereal time) representing the same time instant. East longitudes are
// positive.
func (ct CivilTime) ToLST(longitude float64) SiderealTime {
	return ct.ToGST().ToLST(longitude)
}

// A SiderealTime represents an instant in time defined relative to the position
// of fixed stars without correction for Earth nutation (mean sidereal time).
type SiderealTime struct {
	longitude float64
	timeInSec float64
	date      time.Time
}

// NewGreenwichMeanSiderealTime creates new SiderealTime at 0 longitude
// (Greenwich) corresponding to the given time in proleptic Gregorian calendar
// with accuracy ± 1 second (due to using UTC time instead of UT1).
func NewGreenwichMeanSiderealTime(t time.Time) SiderealTime {
	return NewCivilTime(t).ToGST()
}

// NewLocalMeanSiderealTime creates new SiderealTime at given longitude
// corresponding to the given time in proleptic Gregorian calendar
// with accuracy ± 1 second (due to using UTC time instead of UT1). East
// longitudes are positive.
func NewLocalMeanSiderealTime(t time.Time, longitude float64) SiderealTime {
	return NewCivilTime(t).ToLST(longitude)
}

// String returns the time formatted similar to ISO 8601.
func (st SiderealTime) String() string {
	hour := math.Trunc(st.timeInSec / (60 * 60))
	min := math.Abs(math.Mod(math.Trunc(st.timeInSec/60), 60))
	sec := math.Abs(math.Mod(st.timeInSec, 60))
	return fmt.Sprintf("%v %02.0f:%02.0f:%04.1f %+.4f° GST", st.date.Format(time.DateOnly), hour, min, sec, st.longitude)
}

// ToLST returns SiderealTime at given longitude (local mean sidereal time)
// representing the same time instant. East longitudes are positive.
//
// Method is based on algorithm from J.L. Lawrence "Celestial Calculations"
// (2019). This implementation computes local mean sidereal time without
// the intermediate step of conversion to GST.
func (st SiderealTime) ToLST(longitude float64) SiderealTime {
	if st.longitude == longitude {
		return st
	}

	const hourInSec = 3600
	const dayInSec = 86400
	LST := st.timeInSec - hourInSec*(st.longitude-longitude)/15
	if LST < 0 {
		LST += dayInSec
	}
	if LST > dayInSec {
		LST -= dayInSec
	}

	return SiderealTime{longitude: longitude, timeInSec: LST, date: st.date}
}

// ToGST returns SiderealTime at 0 longitude (Greenwich mean sidereal time)
// representing the same time instant.
func (st SiderealTime) ToGST() SiderealTime {
	if st.longitude == 0 {
		return st
	}
	return st.ToLST(0)
}