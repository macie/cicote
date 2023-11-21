package main

import (
	"fmt"
	"math"
	"time"
)

// sunLocation returns ecliptic coordinates of sun for given UTC time.
func sunLocation(utc time.Time) coordEcliptic {
	return coordEcliptic{}
}

func solarDayBounds(observerLoc coordTerrestial, day time.Time) (err error, rise, set time.Time) {
	sunMidnightLoc := sunLocation(day)
	err, rising_1, setting_1 := starDayBounds(sunMidnightLoc, observerLoc)
	if err != nil {
		return err, rise, set // sun is not rising or setting on given day for observer
	}

	lambda_2 := sunMidnightLoc.lambda + 0.985647 // next midnight location = sun location + sun daily move
	if lambda_2 > 360 {
		lambda_2 -= 360
	}
	sunNextMidnightLoc := coordEcliptic{lambda: lambda_2, beta: sunMidnightLoc.beta}
	err, rising_2, setting_2 := starDayBounds(sunNextMidnightLoc, observerLoc)
	if err != nil {
		return err, rise, set // sun is not rising or setting on given day for observer
	}

	rise = interpolateSiderealTime(rising_1, rising_2).ToLCT()
	set = interpolateSiderealTime(setting_1, setting_2).ToLCT()

	return nil, rise, set
}

func starDayBounds(starLocation coordEcliptic, observerLocation coordTerrestial) (err error, rise, set time.Time) {
	alpha, delta := starLocation.ToEquatorial()
	phi := observerLocation.lat
	H_1 := math.Tan(phi) * math.Tan(delta)
	if math.Abs(H_1) > 1 {
		return fmt.Errorf("star at %v is not rising or not setting for Earth observer at %v", starLocation, observerLocation), time.Time{}, time.Time{}
	}
	H_2 := math.Acos(-H_1) / 15
	LST_r := 24 + alpha - H_2
	if LST_r > 24 {
		LST_r -= 24
	}
	LST_s := alpha + H_2
	if LST_s > 24 {
		LST_s -= 24
	}
	return nil, LST_r, LST_s
}

//
// Accuracy: 0.01
func interpolateSiderealTime(first, second LocalSiderealTime) LocalSiderealTime {
	return (24.07 * first) / (24.07 + first - second)
}

type coordEcliptic struct {
	lambda Angle // longitude
	beta   Angle // latitude
}

type coordEquatorial struct {
	alpha Angle // longitude
	delta Angle // latitude
}