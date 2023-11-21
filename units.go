package main

import "math"

// // DMS is angle measurement unit.
// type DMS struct {
// 	isNegative bool
// 	deg        uint8
// 	min        uint8
// 	sec        uint8
// }

// // NewDMS creates DMS unit from decimal degree.
// func NewDMS(angle float64) DMS {
// 	dms := DMS{}
// 	if angle < 0 {
// 		dms.isNegative = true
// 	}
// 	dec := math.Abs(angle)
// 	full_deg, frac_deg := math.Modf(dec)
// 	dms.deg = uint8(full_deg)
// 	full_min, frac_min := math.Modf(60 * frac_deg)
// 	dms.min = uint8(full_min)
// 	dms.sec = uint8(60 * frac_min)
// 	return dms
// }

// func (d DMS) ToDecimal() float64 {
// 	min := float64(d.sec)/60.0 + float64(d.min)
// 	dec_deg := min / 60.0
// 	total_dec := float64(d.deg) + dec_deg
// 	if d.isNegative {
// 		total_dec = -total_dec
// 	}
// 	return total_dec
// }

type Equatorial struct {
	delta float64 // declination
	alpha float64 // right ascension
}

//func (e Equatorial) HourAngle(lst float32) float32 {
//	h := HMStoDecimal(lst) - HMStoDecimal(e.alpha)
//	if h < 0 {
//		h += 24
//	}
//	return DecimalToHMS(h)
//}

//func (e Equatorial) ToHorizon() Horizon {
//	return Horizon{}
//}

//type Horizon struct {
//	a float32 // azimuth
//	h float32 // altitude
//}

//type Ecliptic struct {
//	beta   float32 // latitude
//	lambda float32 // longitude
//}

//func (e Ecliptic) ToEquatorial(t DateTime) Equatorial {
//	eq := Equatorial{}
//	obliquity := OBLIQUITY_J2000_IN_DEG - (T*(46.815+T*(0.0006-T*0.00181)))/3600
// ...
//	return eq
//}