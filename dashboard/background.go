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
		SPEED_FONT,
		SPEED_LABEL_COLOR,
		SPEED_LINE_HEIGHT,
		LEFT_MARGIN+LABEL_COLUMN,
		TOP_MARGIN,
	)
	d.display.Write("Speed:")
}
