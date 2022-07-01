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
		SPEED_FONT,
		SPEED_DATA_COLOR,
		SPEED_LINE_HEIGHT,
		LEFT_MARGIN+DATA_COLUMN,
		TOP_MARGIN,
	)
	d.display.Write(fmt.Sprintf("%5.2f", d.speed))
}

func (d *dashboardDisplay) printDistance(distance float64) {
	if math.Abs(d.distance-distance) < DISTANCE_RESOLUTION {
		return
	}
	d.distance = distance
	d.setFont(
		DISTANCE_FONT,
		DISTANCE_LABEL_COLOR,
		DISTANCE_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT,
	)
	var dist string
	if distance < 1000 {
		dist = fmt.Sprintf("%4d m", int(d.distance))
	} else {
		dist = fmt.Sprintf("%5.2f km", distance/1000)
	}
	d.display.Write(dist)
}

func (d *dashboardDisplay) printDuration(dur time.Duration) {
	if dur-d.duration < DURATION_RESOLUTION {
		return
	}
	d.duration = dur
	sec := int(d.duration.Seconds()) % 60
	min := sec / 60 % 60
	hour := sec / 3600
	d.setFont(
		DURATION_FONT,
		DURATION_LABEL_COLOR,
		DURATION_LINE_HEIGHT,
		LEFT_MARGIN+DATA_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT+DISTANCE_LINE_HEIGHT,
	)
	d.display.Write(fmt.Sprintf("%02d:%02d:%02d", hour, min, sec))
}
