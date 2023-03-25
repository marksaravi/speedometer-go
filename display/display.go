package display

import (
	"log"
	"time"
)

type display struct {
	resetChannel chan bool
}

func NewDisplay() *display {
	resetChannel := make(chan bool)
	return &display{
		resetChannel: resetChannel,
	}
}

func (d *display) Initialize() {}

func (d *display) Update(speed float64, distance float64, duration time.Duration) {}

func (d *display) Touched(x, y float64) {
	log.Printf("TOUCH: %f,%f\n", x, y)
}

func (d *display) ResetChannel() <-chan bool {
	return d.resetChannel
}
