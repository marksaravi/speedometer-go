package dashboard

import (
	"fmt"
	"time"

	"github.com/marksaravi/devices-go/devices/display"
)

type dashboardDisplay struct {
	display   display.RGBDisplay
	prevSpeed float64
}

func NewDashboardDisplay(display display.RGBDisplay) (chan<- time.Time, chan<- bool) {
	d := dashboardDisplay{
		display:   display,
		prevSpeed: 0,
	}
	return d.start()
}

func (d *dashboardDisplay) start() (chan<- time.Time, chan<- bool) {
	var pulse chan time.Time = make(chan time.Time, 5)
	var reset chan bool = make(chan bool, 1)
	d.initBackground()
	// go d.displayUpdater(pulse, reset)
	return pulse, reset
}

func (d *dashboardDisplay) displayUpdater(pulse <-chan time.Time, reset <-chan bool) {
	for {
		select {
		case t := <-pulse:
			fmt.Println(t)
			d.calcSpeed(t)
			d.incDistance()
			d.updateDuration()
		case _ = <-reset:
			d.reset()
		}
	}
}

func (d *dashboardDisplay) calcSpeed(t time.Time) {

}

func (d *dashboardDisplay) incDistance() {
}

func (d *dashboardDisplay) updateDuration() {

}

func (d *dashboardDisplay) calcAverageSpeed() {

}

func (d *dashboardDisplay) reset() {

}
