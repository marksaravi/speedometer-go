package dashboard

import (
	"math"

	"github.com/marksaravi/devices-go/colors"
)

func (d *dashboardDisplay) drwaSpeedsBackground() {
	// d.display.SetColor(SPEED_CURVE_DASH)
	// d.display.Arc(SPEED_CURVE_CENTER_X, SPEED_CURVE_CENTER_Y, SPEED_CURVE_RADIUS+SPEED_CURVE_WIDTH/2, SPEED_CURVE_START_ANGLE, SPEED_CURVE_END_ANGLE, 1)
	// d.drawSpeedCurve()
}

func (d *dashboardDisplay) drawSpeedCurve() {
	for angle := SPEED_CURVE_START_ANGLE; angle <= SPEED_CURVE_END_ANGLE; angle += SPEED_CURVE_RESOLUTION {
		d.drawSpeedCurveResolution(angle)
	}
}

func (d *dashboardDisplay) drawSpeedCurveResolution(angle float64) {
	r := SPEED_CURVE_RADIUS - SPEED_CURVE_WIDTH/2
	x1 := SPEED_CURVE_CENTER_X + (r)*math.Cos(angle)
	x2 := SPEED_CURVE_CENTER_X + (r-SPEED_CURVE_DASH_LEN)*math.Cos(angle)
	y1 := SPEED_CURVE_CENTER_Y + (r)*math.Sin(angle)
	y2 := SPEED_CURVE_CENTER_Y + (r-SPEED_CURVE_DASH_LEN)*math.Sin(angle)
	d.display.Line(x1, y1, x2, y2)
}

func (d *dashboardDisplay) drawSpeed(speed float64) {
	if speed == d.prevSpeed {
		return
	}

	dSpeed := MAX_SPEED / 1800
	d.prevSpeed = speed - dSpeed*3
	if speed < d.prevSpeed {
		dSpeed = -dSpeed
	}

	r1 := SPEED_CURVE_RADIUS + SPEED_CURVE_WIDTH/2 - 1
	r2 := r1 - SPEED_CURVE_WIDTH + 2
	var color colors.Color = BACKGROUNG_COLOR
	// d.display.Arc(SPEED_CURVE_CENTER_X, SPEED_CURVE_CENTER_Y, SPEED_CURVE_RADIUS+SPEED_CURVE_WIDTH/2-1, toRad(starAngle), toRad(endAngle), SPEED_CURVE_WIDTH-2)
	for s := d.prevSpeed; s < speed; s += dSpeed {
		angle := speedToAngle(s)
		x1 := SPEED_CURVE_CENTER_X + r1*math.Cos(angle)
		x2 := SPEED_CURVE_CENTER_X + r2*math.Cos(angle)
		y1 := SPEED_CURVE_CENTER_Y + r1*math.Sin(angle)
		y2 := SPEED_CURVE_CENTER_Y + r2*math.Sin(angle)
		if dSpeed < 0 {
			color = BACKGROUNG_COLOR
		} else {
			color = speedToColor(s)
		}
		d.display.SetColor(color)
		d.display.Line(x1, y1, x2, y2)
	}
	d.prevSpeed = speed
}
