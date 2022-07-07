package dashboard

import (
	"fmt"
)

const DIGIT_WIDTH int = 36
const COLON_WIDTH int = 14

func (d *dashboardDisplay) printSpeed(speed float64) {
	d.speed = speed
	x := LEFT_MARGIN + DATA_COLUMN
	y := 0
	d.setFont(
		SPEED_DATA_FONT,
		d.theme.SpeedDataColor,
		LEFT_MARGIN+DATA_COLUMN,
		0,
	)
	d.display.ClearArea(float64(x), float64(y+30), float64(x+115), float64(y+SPEED_LABEL_LINE_Y+2))
	d.display.Write(fmt.Sprintf("%3.1f", d.speed))
	// d.display.SetColor(colors.RED)
	// d.display.Rectangle(float64(x), float64(y+30), float64(x+115), float64(y+SPEED_LABEL_LINE_Y+2))
}

func (d *dashboardDisplay) printDistance(distance float64) {
	d.distance = distance
	x := LEFT_MARGIN + DATA_COLUMN
	y := SPEED_LABEL_LINE_Y
	d.setFont(
		DISTANCE_DATA_FONT,
		d.theme.DistanceDataColor,
		x,
		y,
	)
	d.display.ClearArea(float64(x), float64(y+38), float64(x+110), float64(y+DISTANCE_LABEL_LINE_Y))
	d.display.Write(fmt.Sprintf("%4.2f", distance/1000))
	// d.display.SetColor(colors.RED)
	// d.display.Rectangle(float64(x), float64(y+38), float64(x+110), float64(y+DISTANCE_LABEL_LINE_Y))
}

func (d *dashboardDisplay) printDurationDigits(t int, change TimeChanged) {
	drawDigit := func(t, x, y int) {
		d.display.ClearArea(float64(x), float64(y+40), float64(x+DIGIT_WIDTH), float64(y+DURATION_LABEL_LINE_Y))
		d.display.MoveCursor(x, y)
		d.display.Write(fmt.Sprintf("%02d", t))
		// d.display.SetColor(colors.RED)
		// d.display.Rectangle(float64(x), float64(y+40), float64(x+DIGIT_WIDTH), float64(y+DURATION_LABEL_LINE_Y))
	}
	x := LEFT_MARGIN + DATA_COLUMN
	y := SPEED_LABEL_LINE_Y + DISTANCE_LABEL_LINE_Y
	d.setFont(
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		x,
		y,
	)
	if change == HOUR_CHANGED {
		drawDigit(t, x, y)
	}
	x += DIGIT_WIDTH + COLON_WIDTH
	if change == MINUTE_CHANGED {
		drawDigit(t, x, y)
	}
	x += DIGIT_WIDTH + COLON_WIDTH
	if change == SECOND_CHANGED {
		drawDigit(t, x, y)
	}
}
