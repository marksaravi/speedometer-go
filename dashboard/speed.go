package dashboard

import "github.com/marksaravi/devices-go/colors/rgb565"

func (d *dashboardDisplay) drwaSpeedsBackground() {
	d.display.SetColor(rgb565.WHITE)
	d.display.Arc(SPEED_CURVE_CENTER_X, SPEED_CURVE_CENTER_Y, SPEED_CURVE_CENTER_RADIUS, toDeg(SPEED_CURVE_START_ANGLE), toDeg(SPEED_CURVE_END_ANGLE), SPEED_CURVE_CENTER_WIDTH)
}
