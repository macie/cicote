package main

type EquatorialCoords struct {
	declinationInSec    float64 // delta
	rightAscensionInSec float64 // alpha
}

func (e EquatorialCoords) HourAngle(lst SiderealTime) float64 {
	h := lst.timeInSec - e.rightAscensionInSec
	if h < 0 {
		h += 24 * 60 * 60
	}
	return h
}

// func (e Equatorial) ToHorizon() Horizon {
// 	return Horizon{}
// }

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