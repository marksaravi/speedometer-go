package pulsesensor

import (
	"time"

	"github.com/marksaravi/drivers-go/hardware/gpio"
)

type pulseSensor struct {
	lastRead   time.Time
	isHigh     bool
	pulsePinIn gpio.GPIOPinIn
}

func NewPulseSensor(pulsePinIn gpio.GPIOPinIn) *pulseSensor {
	return &pulseSensor{
		pulsePinIn: pulsePinIn,
		lastRead:   time.Now().Add(-time.Second * 86400),
	}
}

func (s *pulseSensor) Read() (bool, time.Duration) {
	pulsed := false
	dur := time.Since(s.lastRead)
	isHigh := s.pulsePinIn.Read()
	if isHigh != s.isHigh {
		if isHigh {
			pulsed = true
		}
		s.isHigh = isHigh
		s.lastRead = time.Now()
	}
	return pulsed, dur
}
