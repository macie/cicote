package main

import (
	"testing"
	"time"
)

func TestDecimalDay(t *testing.T) {
	testcases := map[time.Time]float64{
		time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC): 13,
		time.Date(1, 2, 14, 6, 0, 0, 0, time.UTC):    14.25,
		time.Date(1, 3, 21, 14, 33, 36, 0, time.UTC): 21.606667,
	}
	for tc, want := range testcases {
		if got := round(decimalDay(tc), 6); got != want {
			t.Errorf("decimalDay(\"%v\") = %v, want %v", tc, got, want)
		}
	}
}

func TestDecimalHour(t *testing.T) {
	testcases := map[time.Time]float64{
		time.Date(1, 1, 1, 2, 3, 41, 0, time.UTC):   2.061389,
		time.Date(1, 1, 1, 10, 25, 11, 0, time.UTC): 10.419722,
		time.Date(1, 1, 1, 18, 0, 0, 0, time.UTC):   18,
		time.Date(1, 1, 1, 23, 30, 0, 0, time.UTC):  23.5,
	}
	for tc, want := range testcases {
		if got := round(decimalHour(tc), 6); got != want {
			t.Errorf("decimalHour(\"%v\") = %v, want %v", tc, got, want)
		}
	}
}
