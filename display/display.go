package display

import (
	"fmt"
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
	sketcher.SetRotation(drawings.ROTATION_270)
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
	// go func() {
	// 	s1 := rand.NewSource(time.Now().UnixNano())
	// 	r1 := rand.New(s1)
	// 	for {
	// 		d.writeSpeed(r1.Float64() * 100)
	// 		d.sketcher.Update()
	// 		time.Sleep(time.Second)
	// 	}
	// }()
}

func (d *display) Update(speed float64, distance float64, duration time.Duration) {}

func (d *display) Touched(x, y float64) {
	log.Printf("TOUCH: %f,%f\n", x, y)
}

func (d *display) ResetChannel() <-chan bool {
	return d.resetChannel
}

func (d *display) write(text string, font fonts.BitmapFont, color colors.Color, x, y int, xscale, yscale int) {
	d.sketcher.SetFont(font)
	x1, y1, x2, y2 := d.sketcher.GetTextArea(x, y, text, xscale, yscale)
	d.sketcher.ClearArea(float64(x1), float64(y1), float64(x2), float64(y2), d.theme.BackgroungColor)
	d.sketcher.MoveCursor(x, y)
	d.sketcher.Write(text, color)
}

func (d *display) writeSpeed(speed float64) {
	const x float64 = 20
	const y float64 = 120
	const xScale = 1
	const yScale = 1

	d.write(fmt.Sprintf("%0.1f", speed), fonts.FreeSans24pt7b, d.theme.SpeedColor, int(d.xs+x), int(d.ys+y), xScale, yScale)
}
