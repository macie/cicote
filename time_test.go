package main

import (
	"testing"
	"time"
)

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
		// to the Astronomical Almanac" (2012)
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
	}
	for time, want := range testcases {
		if got := round(NewCivilTime(time).JulianDay(), 7); got != want {
			t.Errorf("NewCivilTime(%v).JulianDay() = %v, want %v", time, got, want)
		}
	}
}
