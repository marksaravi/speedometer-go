package display

import (
	"log"
	"time"

	"github.com/marksaravi/drawings-go/drawings"
	"github.com/marksaravi/drivers-go/colors"
)

type display struct {
	resetChannel chan bool
	sketcher     drawings.Sketcher
}

func NewDisplay(sketcher drawings.Sketcher) *display {
	resetChannel := make(chan bool)
	return &display{
		resetChannel: resetChannel,
		sketcher:     sketcher,
	}
}

func (d *display) Initialize() {
	d.sketcher.Clear(colors.WHITE)
	d.sketcher.FillCircle(160, 120, 50, colors.RED)
	d.sketcher.Update()
}

func (d *display) Update(speed float64, distance float64, duration time.Duration) {}

func (d *display) Touched(x, y float64) {
	log.Printf("TOUCH: %f,%f\n", x, y)
}

func (d *display) ResetChannel() <-chan bool {
	return d.resetChannel
}
