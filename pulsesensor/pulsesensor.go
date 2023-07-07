package pulsesensor

import (
	"time"
	"github.com/marksaravi/drivers-go/hardware/gpio"
)

type pulseSensor struct {
	pulseInput gpio.GPIOPinIn
}

func NewPulseSensor(pulseInput gpio.GPIOPinIn) *pulseSensor {
	return &pulseSensor {
		pulseInput: pulseInput,
	}
}


const MID_DUR time.Duration = time.Second/30
const MAX_DUR = MID_DUR + MID_DUR/4
const MIN_DUR = MID_DUR - MID_DUR/4
var  lastRead time.Time= time.Now()
var dur = MID_DUR
var dt time.Duration = MID_DUR/400

func mock() bool {
	if dur >= MAX_DUR || dur <= MIN_DUR {
		dt = -dt
	}
	if time.Since(lastRead) >= dur {
		lastRead = time.Now()
		dur += dt
		return true
	}
	return false
}

func (s *pulseSensor) Read() bool {
	return mock()
}