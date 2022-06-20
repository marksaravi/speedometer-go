package dashboard

import (
	"fmt"
	"math"
)

func (d *dashboardDisplay) drwaSpeedsBackground() {
	d.display.SetColor(SPEED_CURVE_DASH)
	d.display.Arc(SPEED_CURVE_CENTER_X, SPEED_CURVE_CENTER_Y, SPEED_CURVE_RADIUS+SPEED_CURVE_WIDTH/2, toRad(SPEED_CURVE_START_ANGLE), toRad(SPEED_CURVE_END_ANGLE), 1)
	d.drawSpeedCurve()
}

func (d *dashboardDisplay) drawSpeedCurve() {
	for angle := SPEED_CURVE_START_ANGLE; angle <= SPEED_CURVE_END_ANGLE; angle += SPEED_CURVE_RESOLUTION {
		d.drawSpeedCurveResolution(toRad(angle))
	}
}

func (d *dashboardDisplay) drawSpeedCurveResolution(angle float64) {
	x1 := SPEED_CURVE_CENTER_X + (SPEED_CURVE_RADIUS-SPEED_CURVE_WIDTH/2)*math.Cos(angle)
	x2 := SPEED_CURVE_CENTER_X + (SPEED_CURVE_RADIUS-SPEED_CURVE_WIDTH/2-SPEED_CURVE_DASH_LEN)*math.Cos(angle)
	y1 := SPEED_CURVE_CENTER_Y + (SPEED_CURVE_RADIUS-SPEED_CURVE_WIDTH/2)*math.Sin(angle)
	y2 := SPEED_CURVE_CENTER_Y + (SPEED_CURVE_RADIUS-SPEED_CURVE_WIDTH/2-SPEED_CURVE_DASH_LEN)*math.Sin(angle)
	d.display.Line(x1, y1, x2, y2)
}
func (d *dashboardDisplay) drawSpeed(speed float64) {
	if speed == d.prevSpeed {
		return
	}
	color := SPEED_FORWARD
	if speed < d.prevSpeed {
		fmt.Println("backward", d.prevSpeed, speed)
		color = SPEED_BACKWARD
	} else {
		fmt.Println("forward", d.prevSpeed, speed)
	}

	starAngle := speedToAngle(d.prevSpeed)
	endAngle := speedToAngle(speed)
	if starAngle > endAngle {
		t := starAngle
		starAngle = endAngle
		endAngle = t
	}
	d.display.SetColor(color)
	d.display.Arc(SPEED_CURVE_CENTER_X, SPEED_CURVE_CENTER_Y, SPEED_CURVE_RADIUS+SPEED_CURVE_WIDTH/2-1, toRad(starAngle), toRad(endAngle), SPEED_CURVE_WIDTH-2)
	d.prevSpeed = speed
}
