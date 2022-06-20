package dashboard

import "math"

func toRad(r float64) float64 {
	return r * math.Pi / 180
}

func speedToAngle(speed float64) float64 {
	return SPEED_CURVE_START_ANGLE + (SPEED_CURVE_END_ANGLE-SPEED_CURVE_START_ANGLE)*speed/MAX_SPEED
}
