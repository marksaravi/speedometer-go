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

const DUR = time.Millisecond * 100

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

// func (s *pulseSensor) readPulse() bool {
// 	pulsed := false
// 	level := s.pulsePinIn.Read()
// 	if s.prevPulseLevel != level && level == gpio.Low {
// 		s.pulseCounter++
// 		pulsed = true
// 		s.lastPulseReadTime = time.Now()
// 	}
// 	s.prevPulseLevel = level
// 	return pulsed
// }
