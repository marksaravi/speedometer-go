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
	input gpio.PinIn
	reset gpio.PinIn
	lcd   lcdDisplay

	distPerPulse      float64
	sleepAfterPulseMS int
	counter           int64
	startTime         time.Time

	pulse gpio.Level

	resetLevel gpio.Level
	resetTime  time.Time

	speedPulses []time.Time
	dur         time.Duration
	distance    float64
	speed       float64
}
