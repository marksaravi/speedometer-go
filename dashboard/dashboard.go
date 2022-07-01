package dashboard

import (
	"time"

	"github.com/marksaravi/devices-go/devices/display"
)

type dashboardDisplay struct {
	display  display.RGBDisplay
	speed    float64
	distance float64
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
func (d *dashboardDisplay) Update(speed, distanceKm float64, duration time.Duration) {}
