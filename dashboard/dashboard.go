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

// func (d *dashboardDisplay) start() (chan<- time.Time, chan<- bool) {
// 	var pulse chan time.Time = make(chan time.Time, 5)
// 	var reset chan bool = make(chan bool, 1)

// 	// go d.displayUpdater(pulse, reset)
// 	return pulse, reset
// }

// func (d *dashboardDisplay) calcSpeed(t time.Time) {

// }

// func (d *dashboardDisplay) incDistance() {
// }

// func (d *dashboardDisplay) updateDuration() {

// }

// func (d *dashboardDisplay) calcAverageSpeed() {

// }

// func (d *dashboardDisplay) reset() {

// }
