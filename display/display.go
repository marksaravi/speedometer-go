package display

import (
	"fmt"
	"log"

	"github.com/marksaravi/drawings-go/drawings"
	"github.com/marksaravi/drivers-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
	"github.com/marksaravi/speedometer-go/themes"
)

type area struct {
	x1, x2, y1, y2 float64
}

type display struct {
	resetChannel          chan bool
	sketcher              drawings.Sketcher
	theme                 themes.Theme
	xs, ys, width, height float64
	speedArea             area
	distanceArea          area
	timeArea              area
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
		speedArea:    area{x1: 1000000, y1: 1000000, x2: -1000000, y2: -1000000},
		distanceArea: area{x1: 1000000, y1: 1000000, x2: -1000000, y2: -1000000},
		timeArea:     area{x1: 1000000, y1: 1000000, x2: -1000000, y2: -1000000},
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

func (d *display) Touched(x, y float64) {
	log.Printf("TOUCH: %f,%f\n", x, y)
}

func (d *display) ResetChannel() <-chan bool {
	return d.resetChannel
}

func adjustArea(x1, y1, x2, y2 float64, a *area) {
	if x1 < a.x1 {
		a.x1 = x1
	}
	if y1 < a.y1 {
		a.y1 = y1
	}
	if x2 > a.x2 {
		a.x2 = x2
	}
	if y2 > a.y2 {
		a.y2 = y2
	}
}

func (d *display) write(text string, font fonts.BitmapFont, color colors.Color, x, y float64, xscale, yscale float64) {
	d.sketcher.SetFont(font)
	x1, y1, x2, y2 := d.sketcher.GetTextArea(x, y, text, xscale, yscale)
	adjustArea(x1, y1, x2, y2, &d.speedArea)
	fmt.Println(d.speedArea)
	d.sketcher.ClearArea(d.speedArea.x1, d.speedArea.y1, d.speedArea.x2, d.speedArea.y2, d.theme.BackgroungColor)
	d.sketcher.Rectangle(d.speedArea.x1, d.speedArea.y1, d.speedArea.x2, d.speedArea.y2, colors.YELLOW)
	d.sketcher.MoveCursor(x, y)
	d.sketcher.WriteScaled(text, xscale, yscale, color)
}

func (d *display) writeSpeed(speed float64) {
	const x float64 = 20
	const y float64 = 120
	const xScale = 2
	const yScale = 3

	d.write(fmt.Sprintf("%0.1f", speed), fonts.FreeSans24pt7b, d.theme.SpeedColor, d.xs+x, d.ys+y, xScale, yScale)
}
