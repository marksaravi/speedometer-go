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

func (d *dashboardDisplay) Update(data DisplayData) {
	d.printSpeed(data.Speed, data.SpeedChanged)
	d.printDistance(data.Distance, data.DistanceChanged)
	if data.SecChanged {
		d.printDurationDigits(data.Sec, SEC_CHANGED)
	}
	if data.MinChanged {
		d.printDurationDigits(data.Min, MIN_CHANGED)
	}
	if data.HourChanged {
		d.printDurationDigits(data.Hour, HOUR_CHANGED)
	}
	d.display.Update()
}
