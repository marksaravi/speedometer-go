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

func (d *dashboardDisplay) Update(speed, distance float64, sec, min, hour int, minChanged, hourChanged bool) {
	d.printSpeed(speed)
	d.printDistance(distance)

	d.printDurationDigits(sec, SEC_CHANGED)
	if minChanged {
		d.printDurationDigits(min, MIN_CHANGED)
	}
	if hourChanged {
		d.printDurationDigits(min, HOUR_CHANGED)
	}
	d.display.Update()
}
