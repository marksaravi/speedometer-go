package dashboard

import (
	"fmt"

	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

func (d *dashboardDisplay) printDigits(digits string, font fonts.BitmapFont, color colors.Color, x, y int) {
	x1, y1, x2, y2 := d.display.GetTextArea(digits)
	d.display.ClearArea(float64(x+x1), float64(y+y1), float64(x+x2), float64(y+y2))
	d.setTextSettings(font, color, x, y)
	d.display.Write(digits)
}

func (d *dashboardDisplay) printSpeed(speed float64) {
	d.speed = speed
	x := DATA_X
	y := SPEED_DATA_LINE_Y
	d.printDigits(fmt.Sprintf("%3.1f", d.speed), SPEED_DATA_FONT, d.theme.SpeedDataColor, x, y)
}

func (d *dashboardDisplay) printDistance(distance float64) {
	d.distance = distance
	x := DATA_X
	y := DISTANCE_DATA_LINE_Y
	d.printDigits(fmt.Sprintf("%4.2f", distance/1000), DISTANCE_DATA_FONT, d.theme.DistanceDataColor, x, y)
}

func (d *dashboardDisplay) printDurationDigits(t int, change TimeChanged) {
	digits := fmt.Sprintf("%02d", t)
	x := LABEL_X + TIME_DATA_COLUMN
	y := DURATION_DATA_LINE_Y
	DIGIT_OFFSET := TIME_DIGIT_WIDTH + TIME_COLON_WIDTH
	if change == MINUTE_CHANGED {
		x += DIGIT_OFFSET
	}
	if change == SECOND_CHANGED {
		x += 2 * DIGIT_OFFSET
	}
	d.printDigits(digits, DURATION_DATA_FONT, d.theme.DurationDataColor, x, y)
}

func (d *dashboardDisplay) printDurationColons() {
	x := LABEL_X + TIME_DATA_COLUMN + TIME_DIGIT_WIDTH + TIME_COLON_OFFSET
	y := DURATION_DATA_LINE_Y
	d.writeText(
		":",
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		x,
		y,
	)
	d.writeText(
		":",
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		x+TIME_DIGIT_WIDTH+TIME_COLON_WIDTH,
		y,
	)
}
