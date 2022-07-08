package speedometer

import (
	"time"

	"periph.io/x/conn/v3/gpio"
)

type Config struct {
	DistancePerPulse  float64 `json:"distance-per-pulse"`
	SleepAfterPulseMS int     `json:"sleep-after-pulse-ms"`
}

type lcdDisplay interface {
	Initialise()
	UpdateSpeed(speed float64)
	UpdateDistance(distance float64)
	UpdateDuration(dur, timeType int)
	UpdateDisplay()
}

type speedometerDev struct {
	pulsePinIn        gpio.PinIn
	resetPinIn        gpio.PinIn
	lcd               lcdDisplay
	distPerPulse      float64
	pulseCounter      int64
	displayUpdateTurn int
	prevPulseLevel    gpio.Level
	startOfRidingTime time.Time
	resetPressedTime  time.Time
	displayUpdateTime time.Time
	speedPulses       [2]time.Time
}
