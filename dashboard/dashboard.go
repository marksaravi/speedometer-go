package dashboard

import (
	"time"

	"github.com/marksaravi/devices-go/devices/display"
)

type dashboardDisplay struct {
	display  display.RGBDisplay
	speed    float64 // km/hour
	distance float64 // meter
	duration time.Duration
}

func NewDashboardDisplay(display display.RGBDisplay) *dashboardDisplay {
	return &dashboardDisplay{
		display:  display,
		speed:    0,
		distance: 0,
		duration: 0,
	}
}

func (d *dashboardDisplay) Initialise() {
	d.initBackground()
}

func (d *dashboardDisplay) Update(speed, distance float64, duration time.Duration) {
	d.printSpeed(speed)
	d.printDistance(distance)
	d.printDuration(duration)
	d.display.Update()
}
