package dashboard

import (
	"fmt"

	"github.com/marksaravi/devices-go/devices/display"
)

type TimeChanged = int
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

func (d *dashboardDisplay) UpdateSecond(seconds int) {
	d.printDurationDigits(seconds, SECOND_CHANGED)
}

func (d *dashboardDisplay) UpdateMinute(minutes int) {
	d.printDurationDigits(minutes, MINUTE_CHANGED)
}

func (d *dashboardDisplay) UpdateHour(hours int) {
	d.printDurationDigits(hours, HOUR_CHANGED)
}

func (d *dashboardDisplay) UpdateDisplay() {
	d.display.Update()
}
