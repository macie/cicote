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