package display

import (
	"log"
	"time"

	"github.com/marksaravi/drawings-go/drawings"
	"github.com/marksaravi/drivers-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
	"github.com/marksaravi/speedometer-go/themes"
)

type display struct {
	resetChannel          chan bool
	sketcher              drawings.Sketcher
	theme                 themes.Theme
	xs, ys, width, height float64
}

func NewDisplay(theme themes.Theme, sketcher drawings.Sketcher, margin float64) *display {
	resetChannel := make(chan bool)

	return &display{
		theme:        theme,
		resetChannel: resetChannel,
		sketcher:     sketcher,
		xs:           margin,
		ys:           margin,
		width:        float64(sketcher.ScreenWidth()) - 2*margin,
		height:       float64(sketcher.ScreenHeight()) - 2*margin,
	}
}

func (d *display) Initialize() {
	d.sketcher.Clear(colors.BLACK)
	d.sketcher.ClearArea(d.xs, d.ys, d.xs+d.width, d.ys+d.height, d.theme.BackgroungColor)
	d.write("12.3", fonts.FreeSans24pt7b, d.theme.SpeedColor, d.xs+10, d.ys+100)
	d.sketcher.Update()
}

func (d *display) Update(speed float64, distance float64, duration time.Duration) {}

func (d *display) Touched(x, y float64) {
	log.Printf("TOUCH: %f,%f\n", x, y)
}

func (d *display) ResetChannel() <-chan bool {
	return d.resetChannel
}

func (d *display) write(text string, font fonts.BitmapFont, color colors.Color, x, y float64) {
	// x1, y1, x2, y2 := d.sketcher.GetTextArea(text)
	d.sketcher.SetFont(font)
	d.sketcher.MoveCursor(int(x), int(y))
	d.sketcher.Write(text, color)
}
