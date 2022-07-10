package speedometer

import (
	"time"

	"github.com/marksaravi/devices-go/hardware/gpio"
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
	pulsePinIn        gpio.GPIOPinIn
	resetPinIn        gpio.GPIOPinIn
	lcd               lcdDisplay
	distPerPulse      float64
	pulseCounter      int64
	displayUpdateTurn int
	prevPulseLevel    gpio.Level
	startOfRidingTime time.Time
	resetPressedTime  time.Time
	speedPulseFrom    time.Time
	speedPulseTo      time.Time
	speedLastUpdate   time.Time
	distLastUpdate    time.Time
	prevSecond        int
	prevMinute        int
	prevHour          int
}
