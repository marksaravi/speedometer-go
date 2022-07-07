package dashboard

import (
	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

func (d *dashboardDisplay) initBackground() {
	d.display.SetBackgroundColor(d.theme.BackgroungColor)
	d.display.Clear()
	d.printLabels()
	d.printSpeed(0)
	d.printDistance(0)
	d.printDurationColons()
	d.printDurationDigits(0, SECOND_CHANGED)
	d.printDurationDigits(0, MINUTE_CHANGED)
	d.printDurationDigits(0, HOUR_CHANGED)
	d.drawGrids()
	d.display.Update()
}

func (d *dashboardDisplay) setFont(font fonts.BitmapFont, color colors.Color, x, y int) {
	d.display.SetColor(color)
	d.display.SetFont(font)
	d.display.MoveCursor(x, y)
}

func (d *dashboardDisplay) printLabels() {
	d.setFont(
		SPEED_LABEL_FONT,
		d.theme.SpeedLabelColor,
		LEFT_MARGIN+LABEL_COLUMN,
		SPEED_LABEL_LINE_Y,
	)
	d.display.Write("Speed (km/h):")
	d.setFont(
		DISTANCE_LABEL_FONT,
		d.theme.DistanceLabelColor,
		LEFT_MARGIN+LABEL_COLUMN,
		DISTANCE_LABEL_LINE_Y,
	)
	d.display.Write("Distance (km):")
	d.setFont(
		DURATION_LABEL_FONT,
		d.theme.DurationLabelColor,
		LEFT_MARGIN+LABEL_COLUMN,
		DURATION_LABEL_LINE_Y,
	)
	d.display.Write("Duration:")
}

func (d *dashboardDisplay) printDurationColons() {
	drawDigit := func(x, y int) {
		d.display.MoveCursor(x+DIGIT_WIDTH+4, y)
		d.display.Write(":")
	}
	x := LEFT_MARGIN + DATA_COLUMN
	y := DURATION_DATA_LINE_Y
	d.setFont(
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		x,
		y,
	)
	drawDigit(x, y)
	x += DIGIT_WIDTH + COLON_WIDTH
	drawDigit(x, y)
}

func (d *dashboardDisplay) drawGrids() {
	d.display.SetColor(colors.LIGHTGRAY)
	for x := float64(32); x < 319; x += 32 {
		d.display.Line(x, 0, x, 239)
	}
	for y := float64(24); y < 239; y += 24 {
		d.display.Line(0, y, 319, y)
	}
}
