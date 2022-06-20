package dashboard

import "github.com/marksaravi/fonts-go/fonts"

func (d *dashboardDisplay) initBackground() {
	d.display.SetBackgroundColor(BACKGROUNG_COLOR)
	d.display.Clear()
	d.display.SetFont(fonts.FreeMono9pt7b)
	d.display.MoveCursor(250, 10)
	d.display.Write("SPEED")
	d.drwaSpeedsBackground()
	d.display.Update()
}
