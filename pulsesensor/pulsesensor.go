package pulsesensor

import (
	"time"
	"github.com/marksaravi/drivers-go/hardware/gpio"
)

type pulseSensor struct {
	lastRead   time.Time
	pulseInput gpio.GPIOPinIn
}

func NewPulseSensor(pulseInput gpio.GPIOPinIn) *pulseSensor {
	return &pulseSensor {
		pulseInput: pulseInput,
		lastRead: time.Now().Add(-time.Second*86400),
	}
}

const DUR = time.Millisecond * 100

func (s *pulseSensor) Read() (bool, time.Duration) {
	dur := time.Since(s.lastRead)
	if dur >= DUR {
		s.lastRead = time.Now()
		return true, dur 
	}
	return false, dur
}