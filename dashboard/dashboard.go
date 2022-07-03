package dashboard

import (
	"math"

	"github.com/marksaravi/devices-go/devices/display"
)

type TimeChanged = int
type dashboardDisplay struct {
	display display.RGBDisplay
	theme   Theme

	second   int
	minute   int
	hour     int
	speed    float64
	distance float64
}

func NewDashboardDisplay(display display.RGBDisplay) *dashboardDisplay {
	return &dashboardDisplay{
		display:  display,
		theme:    DarkTheme,
		speed:    -1,
		distance: -1,
		second:   -1,
		minute:   -1,
		hour:     -1,
	}
}

func (d *dashboardDisplay) Initialise() {
	d.initBackground()
}

func (d *dashboardDisplay) UpdateSpeed(speed float64) {
	if math.Abs(speed-d.speed) > 0.25 {
		d.printSpeed(speed)
		d.speed = speed
	}

}

func (d *dashboardDisplay) UpdateDistance(distance float64) {
	if distance != d.distance {
		d.printDistance(distance)
		d.distance = distance
	}

}

func (d *dashboardDisplay) UpdateSecond(seconds int) {
	if seconds != d.second {
		d.printDurationDigits(seconds, SECOND_CHANGED)
		d.second = seconds
	}

}

func (d *dashboardDisplay) UpdateMinute(minutes int) {
	if minutes != d.minute {
		d.printDurationDigits(minutes, MINUTE_CHANGED)
		d.minute = minutes
	}

}

func (d *dashboardDisplay) UpdateHour(hours int) {
	if hours != d.hour {
		d.printDurationDigits(hours, HOUR_CHANGED)
		d.hour = hours
	}

}

func (d *dashboardDisplay) UpdateDisplay() {
	d.display.Update()
}
