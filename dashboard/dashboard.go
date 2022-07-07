package dashboard

import (
	"fmt"

	"github.com/marksaravi/devices-go/devices/display"
)

type dashboardDisplay struct {
	display display.RGBDisplay
	theme   Theme
}

func NewDashboardDisplay(display display.RGBDisplay) *dashboardDisplay {
	return &dashboardDisplay{
		display: display,
		theme:   DarkTheme,
	}
}

func (d *dashboardDisplay) Initialise() {
	d.initBackground()
}

func (d *dashboardDisplay) UpdateSpeed(speed float64) {
	x := DATA_X
	y := SPEED_DATA_LINE_Y
	d.printDigits(fmt.Sprintf("%3.1f", speed), SPEED_DATA_FONT, d.theme.SpeedDataColor, x, y)

}

func (d *dashboardDisplay) UpdateDistance(distance float64) {
	x := DATA_X
	y := DISTANCE_DATA_LINE_Y
	d.printDigits(fmt.Sprintf("%4.2f", distance/1000), DISTANCE_DATA_FONT, d.theme.DistanceDataColor, x, y)
}

func (d *dashboardDisplay) UpdateDuration(t int, change int) {
	digits := fmt.Sprintf("%02d", t)
	x := DATA_X
	y := DURATION_DATA_LINE_Y
	DIGIT_OFFSET := TIME_DIGIT_WIDTH + TIME_COLON_WIDTH
	if change == MINUTE_CHANGED {
		x += DIGIT_OFFSET
	}
	if change == SECOND_CHANGED {
		x += 2 * DIGIT_OFFSET
	}
	d.printDigits(digits, DURATION_DATA_FONT, d.theme.DurationDataColor, x, y)
}

func (d *dashboardDisplay) UpdateDisplay() {
	d.display.Update()
}
