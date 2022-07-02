package dashboard

import (
	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

func (d *dashboardDisplay) initBackground() {
	d.display.SetBackgroundColor(d.theme.BackgroungColor)
	d.display.Clear()
	d.printLabels()
	d.printDurationDigits(0, SEC_CHANGED)
	d.printDurationDigits(0, MIN_CHANGED)
	d.printDurationDigits(0, HOUR_CHANGED)
	d.printDurationColons()
	d.display.Update()
}

func (d *dashboardDisplay) setFont(font fonts.BitmapFont, color colors.Color, lineHeight int, x, y int) {
	d.display.SetColor(color)
	d.display.SetFont(font)
	d.display.SetLineHeight(lineHeight)
	d.display.MoveCursor(x, y)
}

func (d *dashboardDisplay) printLabels() {
	d.setFont(
		SPEED_LABEL_FONT,
		d.theme.SpeedLabelColor,
		SPEED_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN,
	)
	d.display.Write("Speed (km/h):")
	d.setFont(
		DISTANCE_LABEL_FONT,
		d.theme.DistanceLabelColor,
		DISTANCE_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT,
	)
	d.display.Write("Distance (km):")
	d.setFont(
		DURATION_LABEL_FONT,
		d.theme.DurationLabelColor,
		DURATION_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT+DISTANCE_LINE_HEIGHT,
	)
	d.display.Write("Duration:")
}

func (d *dashboardDisplay) printDurationColons() {
	drawDigit := func(x, y int) {
		d.display.MoveCursor(x+DIGIT_WIDTH+4, y)
		d.display.Write(":")
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
	drawDigit(x, y)
	x += DIGIT_WIDTH + COLON_WIDTH
	drawDigit(x, y)
}
