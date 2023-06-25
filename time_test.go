package main

import (
	"strconv"
	"testing"
	"time"
)

func TestCivilTimeDate(t *testing.T) {
	testcase := time.Date(2023, 6, 15, 6, 9, 0, 0, time.UTC)
	wantY, wantM, wantD := 2023.0, 6.0, 15.25625
	if gotY, gotM, gotD := NewCivilTime(testcase).Date(); gotY != wantY || gotM != wantM || gotD != wantD {
		t.Errorf("NewCivilTime(%v).Date() = (%v, %v, %v), want (%v, %v, %v)", testcase, gotY, gotM, gotD, wantY, wantM, wantD)
	}
}

func TestCivilTimeEqual(t *testing.T) {
	now := time.Now()
	local := NewCivilTime(now.Local())
	utc := NewCivilTime(now.UTC())
	if !local.Equal(utc) {
		t.Errorf("CivilTime %v is not equal %v", local, utc)
	}
}

func TestJulianDay(t *testing.T) {
	// dates before 1582-10-15 are converted from Julian calendar to Gregorian using:
	// https://legacy-www.math.harvard.edu/computing/javascript/Calendar/index.html
	testcases := map[time.Time]float64{
		// test cases from J. Meeus "Astronomical Algorithms" (1998)
		time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC):    2451545,
		time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC):     2451179.5,
		time.Date(1988, 6, 19, 12, 0, 0, 0, time.UTC):   2447332,
		time.Date(1988, 1, 27, 0, 0, 0, 0, time.UTC):    2447187.5,
		time.Date(1987, 6, 19, 12, 0, 0, 0, time.UTC):   2446966,
		time.Date(1987, 1, 27, 0, 0, 0, 0, time.UTC):    2446822.5,
		time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC):     2415020.5,
		time.Date(1600, 12, 31, 0, 0, 0, 0, time.UTC):   2305812.5,
		time.Date(1600, 1, 1, 0, 0, 0, 0, time.UTC):     2305447.5,
		time.Date(837, 4, 14, 8, 0, 0, 0, time.UTC):     2026871.8333333, // Julian:  0873-04-10
		time.Date(-123, 12, 29, 0, 0, 0, 0, time.UTC):   1676497.5,       // Julian: -0122-01-01
		time.Date(-123, 12, 28, 0, 0, 0, 0, time.UTC):   1676496.5,       // Julian: -0123-12-31
		time.Date(-1000, 7, 2, 12, 0, 0, 0, time.UTC):   1356001,         // Julian: -1000-07-12
		time.Date(-1000, 2, 19, 0, 0, 0, 0, time.UTC):   1355866.5,       // Julian: -1000-02-29
		time.Date(-1001, 8, 7, 21, 36, 0, 0, time.UTC):  1355671.4,       // Julian: -1001-08-17
		time.Date(-4713, 11, 24, 12, 0, 0, 0, time.UTC): 0,               // Julian: -4712-01-01
		// test cases from S. Urban and P.K. Seidelmann "Explanatory Supplement
		// to the Astronomical Almanac" (2013)
		time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC):  2451545,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC):   2451544.5,
		time.Date(1984, 2, 2, 12, 0, 0, 0, time.UTC):  2445733,
		time.Date(1844, 3, 21, 12, 0, 0, 0, time.UTC): 2394647,
		time.Date(1792, 9, 22, 12, 0, 0, 0, time.UTC): 2375840,
		time.Date(622, 7, 19, 12, 0, 0, 0, time.UTC):  1948440, // Julian:  0622-07-16
		time.Date(139, 7, 19, 12, 0, 0, 0, time.UTC):  1772028, // Julian:  0139-07-20
		time.Date(79, 3, 22, 12, 0, 0, 0, time.UTC):   1749995, // Julian:  0079-03-24
		time.Date(0, 12, 30, 12, 0, 0, 0, time.UTC):   1721424, // Julian:  0001-01-01
		time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC):     1721426, // Julian:  0001-01-03
		time.Date(-746, 2, 18, 12, 0, 0, 0, time.UTC): 1448638, // Julian: -0746-02-26
		time.Date(-3760, 9, 7, 12, 0, 0, 0, time.UTC): 347998,  // Julian: -3760-10-07
		// test cases from J.L. Lawrence "Celestial Calculations" (2019)
		time.Date(2015, 5, 10, 18, 0, 0, 0, time.UTC): 2457153.25,
		time.Date(2015, 5, 10, 6, 0, 0, 0, time.UTC):  2457152.75,
		time.Date(2015, 3, 21, 12, 0, 0, 0, time.UTC): 2457103,
		time.Date(2010, 11, 1, 0, 0, 0, 0, time.UTC):  2455501.5,
		time.Date(2010, 2, 7, 0, 0, 0, 0, time.UTC):   2455234.5,
		time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC):   2455197.5,
		time.Date(1985, 2, 7, 0, 0, 0, 0, time.UTC):   2446103.5,
		time.Date(1858, 11, 17, 0, 0, 0, 0, time.UTC): 2400000.5,
		time.Date(1582, 10, 16, 0, 0, 0, 0, time.UTC): 2299161.5,
		time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC): 2299160.5,
	}
	for time, want := range testcases {
		if got := round(NewCivilTime(time).JulianDay(), 7); got != want {
			t.Errorf("NewCivilTime(%v).JulianDay() = %v, want %v", time, got, want)
		}
	}
}

func TestCivilTimeToLCT(t *testing.T) {
	now := time.Now().UTC()
	offset := -2*time.Hour - 30*time.Minute
	want := NewCivilTime(now)
	if got := NewCivilTime(now).ToLCT(offset); !got.Equal(want) {
		t.Errorf("NewCivilTime(%v).ToLCT(%v) is %v, want %v", now, offset, got, want)
	}
}

func TestCivilTimeToUTC(t *testing.T) {
	now := time.Now()
	want := NewCivilTime(now.UTC())
	if got := NewCivilTime(now).ToUTC(); !got.Equal(want) {
		t.Errorf("NewCivilTime(%v).ToUTC() is %v, want %v", now, got, want)
	}
}

func TestCivilTimeToGST(t *testing.T) {
	testcases := map[time.Time]string{
		// test cases from J. Meeus "Astronomical Algorithms" (1998)
		time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC): "1987-04-10 08:34:57.1 +0.0000° GST", // book:  8h34m57.0896s
		time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC):   "1987-04-10 13:10:46.4 +0.0000° GST", // book: 13h10m46.1351s
		// test cases from J.L. Lawrence "Celestial Calculations" (2019)
		time.Date(2010, 12, 13, 1, 0, 0, 0, time.UTC): "2010-12-13 06:26:27.0 +0.0000° GST", // book: 6h26m34s
		time.Date(2010, 2, 7, 23, 30, 0, 0, time.UTC): "2010-02-07 08:41:53.2 +0.0000° GST", // book: 8h41m53s
		time.Date(2000, 7, 5, 7, 0, 0, 0, time.UTC):   "2000-07-05 01:54:20.6 +0.0000° GST", // book: 1h54m20s
	}
	for tc, want := range testcases {
		if got := NewCivilTime(tc).ToGST().String(); got != want {
			t.Errorf("CivilTime(%v).ToGST() is %v, want %v", tc, got, want)
		}
	}
}

func TestSiderealTimeString(t *testing.T) {
	testcases := map[float64]string{
		-(2*3600 + 20*60 + 1): "2025-06-23 -2:20:01.0 +15.2000° GST",
		0:                     "2025-06-23 00:00:00.0 +15.2000° GST",
		1*3600 + 30*60 + 2.45: "2025-06-23 01:30:02.4 +15.2000° GST",
		50*3600 + 3*60 + 59:   "2025-06-23 50:03:59.0 +15.2000° GST",
	}
	for tc, want := range testcases {
		if got := (SiderealTime{date: time.Date(2025, 6, 23, 1, 2, 3, 0, time.UTC), hourAngleInSec: tc, longitude: 15.2}).String(); got != want {
			t.Errorf("SiderealTime(%v).String() = %v, want %v", tc, got, want)
		}
	}
}

func TestSiderealTimeToLST(t *testing.T) {
	testcases := map[[2]string]string{
		// test cases from P. Duffet-Smith and J. Zwart "Practical Astronomy
		// with your Calculator or Spreadsheet" (2011)
		{"1980-04-22T14:36:51.67Z", "0"}: "1980-04-22 04:40:05.2 +0.0000° GST",
		// test cases from J.L. Lawrence "Celestial Calculations" (2019)
		{"2014-12-12T20:00:00-05:00", "-77"}: "2014-12-13 01:18:34.4 -77.0000° GST",
		{"2000-07-05T12:00:00+05:00", "60"}:  "2000-07-05 05:54:20.6 +60.0000° GST",
	}
	for tc, want := range testcases {
		localDate, _ := time.Parse(time.RFC3339, tc[0])
		long, _ := strconv.ParseFloat(tc[1], 64)
		if got := NewGreenwichMeanSiderealTime(localDate).ToLST(long).String(); got != want {
			t.Errorf("NewGreenwichMeanSiderealTime(\"%v\").ToLST(%v) = \"%v\", want \"%v\"", tc[0], long, got, want)
		}
	}
}