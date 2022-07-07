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

func (d *dashboardDisplay) printDurationDigits(t int, change TimeChanged) {
	digits := fmt.Sprintf("%02d", t)
	x := DATA_X
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
	x := DATA_X + TIME_DIGIT_WIDTH + TIME_COLON_OFFSET
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
