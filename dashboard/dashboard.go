package dashboard

import (
	"time"

	"github.com/marksaravi/devices-go/devices/display"
)

type dashboardDisplay struct {
	display display.RGBDisplay
}

func NewDashboardDisplay(display display.RGBDisplay) *dashboardDisplay {
	return &dashboardDisplay{
		display: display,
	}
}

func (d *dashboardDisplay) Initialise() {
	d.initBackground()
}
func (d *dashboardDisplay) Update(speed, distanceKm float64, duration time.Duration) {}
