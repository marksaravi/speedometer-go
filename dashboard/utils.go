package dashboard

import (
	"github.com/marksaravi/devices-go/colors/rgb565"
)

func speedToAngle(speed float64) float64 {
	return SPEED_CURVE_START_ANGLE + (SPEED_CURVE_END_ANGLE-SPEED_CURVE_START_ANGLE)*speed/MAX_SPEED
}

func speedToColor(speed float64) rgb565.RGB565 {
	if speed < MAX_SAFE_SPEED {
		return SPEED_FORWARD_SAFE
	}
	if speed < MAX_WARNING_SPEED {
		return SPEED_FORWARD_WARN
	}
	return SPEED_FORWARD_DANGER
}
