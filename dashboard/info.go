package dashboard

import (
	"fmt"
	"math"
	"time"
)

func (d *dashboardDisplay) printSpeed(speed float64) {
	if math.Abs(d.speed-speed) < SPEED_RESOLUTION {
		return
	}
	d.speed = speed
	d.setFont(
		SPEED_DATA_FONT,
		d.theme.SpeedDataColor,
		SPEED_LINE_HEIGHT,
		LEFT_MARGIN+DATA_COLUMN,
		TOP_MARGIN,
	)
	d.display.Write(fmt.Sprintf("%3.1f", d.speed))
}

func (d *dashboardDisplay) printDistance(distance float64) {
	if math.Abs(d.distance-distance) < DISTANCE_RESOLUTION {
		return
	}
	d.distance = distance
	d.setFont(
		DISTANCE_DATA_FONT,
		d.theme.DistanceDataColor,
		DISTANCE_LINE_HEIGHT,
		LEFT_MARGIN+DATA_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT,
	)
	d.display.Write(fmt.Sprintf("%5.3f", distance/1000))
}

func (d *dashboardDisplay) printDuration(dur time.Duration) {
	sec := int(dur.Seconds()) % 60
	min := sec / 60 % 60
	hour := sec / 3600
	x := LEFT_MARGIN + DATA_COLUMN
	y := TOP_MARGIN + SPEED_LINE_HEIGHT + DISTANCE_LINE_HEIGHT
	d.setFont(
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		DURATION_LINE_HEIGHT,
		x,
		y,
	)
	// d.display.Write(fmt.Sprintf("%02d:%02d:%02d", hour, min, sec))
	d.display.Write(fmt.Sprintf("%02d", hour))
	x += 32
	d.display.MoveCursor(x, y)
	d.display.Write(":")

	x += 16
	d.display.MoveCursor(x, y)
	d.display.Write(fmt.Sprintf("%02d", min))

	x += 16
	d.display.MoveCursor(x, y)
	d.display.Write(":")

	x += 16
	d.display.MoveCursor(x, y)
	d.display.Write(fmt.Sprintf("%02d", sec))
}
