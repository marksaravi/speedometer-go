package dashboard

import (
	"fmt"
)

const DIGIT_WIDTH int = 36
const COLON_WIDTH int = 14

func (d *dashboardDisplay) printSpeed(speed float64, speedChanged bool) {
	if !speedChanged {
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

func (d *dashboardDisplay) printDistance(distance float64, distanceChanged bool) {
	if !distanceChanged {
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

func (d *dashboardDisplay) printDurationDigits(t int, change TimeChanged) {
	drawDigit := func(t, x, y int) {
		d.display.ClearArea(float64(x), float64(y), float64(x+DIGIT_WIDTH), float64(y+DURATION_LINE_HEIGHT))
		d.display.MoveCursor(x, y)
		d.display.Write(fmt.Sprintf("%02d", t))
	}
	x := LEFT_MARGIN + DATA_COLUMN
	y := TOP_MARGIN + SPEED_LINE_HEIGHT + DISTANCE_LINE_HEIGHT
	d.setFont(
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		DURATION_LINE_HEIGHT,
		x,
		y,
	)
	if change == HOUR_CHANGED {
		drawDigit(t, x, y)
	}
	x += DIGIT_WIDTH + COLON_WIDTH
	if change == MIN_CHANGED {
		drawDigit(t, x, y)
	}
	x += DIGIT_WIDTH + COLON_WIDTH
	if change == SEC_CHANGED {
		drawDigit(t, x, y)
	}
}
