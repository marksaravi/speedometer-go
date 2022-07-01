package dashboard

import (
	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

func (d *dashboardDisplay) initBackground() {
	d.display.SetBackgroundColor(BACKGROUNG_COLOR)
	d.display.Clear()
	d.printLabels()
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
		SPEED_LABEL_COLOR,
		SPEED_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN,
	)
	d.display.Write("Speed (km/h):")
	d.setFont(
		DISTANCE_LABEL_FONT,
		DISTANCE_LABEL_COLOR,
		DISTANCE_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT,
	)
	d.display.Write("Distance (km):")
	d.setFont(
		DURATION_LABEL_FONT,
		DURATION_LABEL_COLOR,
		DURATION_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN+SPEED_LINE_HEIGHT+DISTANCE_LINE_HEIGHT,
	)
	d.display.Write("Duration:")
}
