package speedometer

import (
	"time"

	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
)

type Config struct {
	DistancePerPulse  float64 `json:"distance-per-pulse"`
	SleepAfterPulseMS int     `json:"sleep-after-pulse-ms"`
}

type lcdDisplay interface {
	Initialise()
	Update(dashboard.DisplayData)
}

type speedometerDev struct {
	input gpio.PinIn
	reset gpio.PinIn
	lcd   lcdDisplay

	distPerPulse      float64
	sleepAfterPulseMS int
	counter           int64
	startTime         time.Time

	pulse     gpio.Level
	pulseTime time.Time
	pulseDur  time.Duration

	resetLevel gpio.Level
	resetTime  time.Time

	sec      int
	min      int
	hour     int
	distance float64
	speed    float64
}
