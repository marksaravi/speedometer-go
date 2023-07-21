package pulsesensor

import (
	"time"

	"github.com/marksaravi/drivers-go/hardware/gpio"
)

type pulseSensor struct {
	lastRead   time.Time
	wasHigh    bool
	pulsePinIn gpio.GPIOPinIn
}

func NewPulseSensor(pulsePinIn gpio.GPIOPinIn) *pulseSensor {
	return &pulseSensor{
		pulsePinIn: pulsePinIn,
		lastRead:   time.Now().Add(-time.Second * 86400),
	}
}

func (s *pulseSensor) Read() bool {
	pulsed := false
	wasHigh := s.pulsePinIn.Read()
	if wasHigh != s.wasHigh {
		if wasHigh {
			pulsed = true
		}
		s.wasHigh = wasHigh
		s.lastRead = time.Now()
	}
	return pulsed
}
