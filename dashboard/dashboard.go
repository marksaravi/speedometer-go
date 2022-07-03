package dashboard

import (
	"github.com/marksaravi/devices-go/devices/display"
)

type TimeChanged = int
type dashboardDisplay struct {
	display  display.RGBDisplay
	speed    float64 // km/hour
	distance float64 // meter
	theme    Theme
}

func NewDashboardDisplay(display display.RGBDisplay) *dashboardDisplay {
	return &dashboardDisplay{
		display:  display,
		speed:    0,
		distance: 0,
		theme:    DarkTheme,
	}
}

func (d *dashboardDisplay) Initialise() {
	d.initBackground()
}

func (d *dashboardDisplay) UpdateSpeed(speed float64) {
	d.printSpeed(speed)
}

func (d *dashboardDisplay) UpdateDistance(distance float64) {
	d.printDistance(distance)
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
