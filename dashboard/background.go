package dashboard

import (
	"fmt"
	"time"

	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/devices-go/devices/display"
	"github.com/marksaravi/fonts-go/fonts"
)

func (d *dashboardDisplay) initBackground() {
	fmt.Println("func initBackground")
	d.display.SetBackgroundColor(BACKGROUNG_COLOR)
	d.display.Clear()
	d.display.SetColor(colors.WHEAT)
	d.display.MoveCursor(0, 0)
	d.display.SetColor(colors.BLUE)
	d.display.SetLineHeight(18)
	d.display.SetFont(fonts.FreeSansBoldOblique9pt7b)
	d.display.Write("Hello Mark!")

	d.display.MoveCursor(0, 30)
	d.display.SetColor(colors.RED)
	d.display.SetLineHeight(22)
	d.display.SetFont(fonts.FreeSansOblique18pt7b)
	d.display.Write("Hello Mark!")

	d.display.MoveCursor(0, 70)
	d.display.SetColor(colors.BLACK)
	d.display.SetLineHeight(22)
	d.display.SetFont(fonts.FreeSerif18pt7b)
	d.display.Write("Hello Mark!")
	d.display.SetFont(fonts.FreeSans18pt7b)
	d.display.ThickCircle(160, 120, 100, 20, display.INNER_WIDTH)

	d.display.Update()
	time.Sleep(time.Second)
}
