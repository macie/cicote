package main

import (
	"testing"
	"time"
)

func TestSunLocation(t *testing.T) {
	time, _ := time.Parse(time.RFC1123, "2015-02-05")
	want := coordEcliptic{}
	if got := sunLocation(time); got != want {
		t.Errorf("sunLocation(%v) = %v, want %v", time, got, want)
	}
}

func ExampleSunriseAndSunset() {
	obs := earthObserver{lat: 0, long: 0}
	day, _ := time.Parse(time.RFC1123, "2015-02-05")
	dayMidnight := day.Add(-time.Hour * 24)
	midnightLocation := sunLocation(dayMidnight)
	riseTime := obs.rising(midnightLocation)
	setTime := obs.setting(midnightLocation)
	nextDayMidnight := dayMidnight.Add(time.Hour * 24)
}
