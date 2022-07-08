package speedometer

import (
	"fmt"
	"time"

	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
)

const (
	DISPLAY_UPDATE_TIMEOUT_MS = 100
	SPEED_CALC_PERIOD         = time.Second * 5
)

func NewSpeedometer() *speedometerDev {
	config := ReadConfigs()
	lcd := createDisplay()
	lcd.Initialise()

	speedo := speedometerDev{
		pulsePinIn:        createGpioInputPin("GPIO14"),
		resetPinIn:        createGpioInputPin("GPIO15"),
		lcd:               lcd,
		distPerPulse:      config.DistancePerPulse,
		startOfRidingTime: time.Now(),
		resetPressedTime:  time.Now(),
		speedPulses:       getSpeedPulsesZeroValue(),
		pulseCounter:      0,
		displayUpdateTurn: 0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	for {
		if s.readPulse() {
			s.pushSpeedPulse(time.Now())
			s.triggerUpdate()
		}
		if s.canUpdate() {
			s.update()
		}
		s.readReset()
	}
}

func (s *speedometerDev) readReset() {
	if s.resetPinIn.Read() == gpio.Low {
		s.resetPressedTime = time.Now()
	}
	if time.Since(s.resetPressedTime) > time.Second*3 {
		s.reset()
	}
}

func (s *speedometerDev) reset() {
	s.pulseCounter = 0
	s.startOfRidingTime = time.Now()
	s.prevPulseLevel = gpio.Low
	s.speedPulses = getSpeedPulsesZeroValue()
}

func getSpeedPulsesZeroValue() [2]time.Time {
	return [2]time.Time{time.Time{}, time.Time{}}
}

func (s *speedometerDev) readPulse() bool {
	pulsed := false
	level := s.pulsePinIn.Read()
	if s.prevPulseLevel != level && level == gpio.Low {
		s.pulseCounter++
	}
	s.prevPulseLevel = level
	return pulsed
}

func (s *speedometerDev) canUpdate() bool {
	if time.Since(s.displayUpdateTime) < time.Millisecond*DISPLAY_UPDATE_TIMEOUT_MS {
		return false
	}
	s.displayUpdateTime = time.Now()
	return true
}

func (s *speedometerDev) triggerUpdate() {
	const DT = time.Millisecond * DISPLAY_UPDATE_TIMEOUT_MS / 5
	if time.Since(s.displayUpdateTime) > time.Millisecond*DISPLAY_UPDATE_TIMEOUT_MS-DT {
		s.displayUpdateTime = time.Now().Add(-time.Millisecond*DISPLAY_UPDATE_TIMEOUT_MS - DT)
	}
}

func (s *speedometerDev) calcSpeedDistanceDuration() (
	seconds, minutes, hours int, speed, distance float64,
) {

	speed = s.calcSpeed()
	seconds, minutes, hours = getSecMinHour(time.Since(s.startOfRidingTime))
	distance = s.distPerPulse * float64(s.pulseCounter)
	return
}

func getSecMinHour(d time.Duration) (int, int, int) {
	seconds := int(d.Seconds())
	return seconds % 60, seconds / 60 % 60, seconds / 3600
}

func (s *speedometerDev) update() bool {
	ts := time.Now()
	seconds, minutes, hours, speed, distance := s.calcSpeedDistanceDuration()

	switch s.displayUpdateTurn {
	case 0:
		s.lcd.UpdateDuration(seconds, dashboard.SECOND_CHANGED)
	case 1:
		s.lcd.UpdateDuration(minutes, dashboard.MINUTE_CHANGED)
	case 2:
		s.lcd.UpdateDuration(hours, dashboard.HOUR_CHANGED)
	case 3:
		s.lcd.UpdateSpeed(speed)
	case 4:
		s.lcd.UpdateDistance(distance)
	}
	s.displayUpdateTurn++
	if s.displayUpdateTurn == 5 {
		s.displayUpdateTurn = 0
	}
	s.lcd.UpdateDisplay()
	fmt.Println(time.Since(ts))
	return true
}

func (s *speedometerDev) pushSpeedPulse(t time.Time) {
	if s.speedPulses[1].IsZero() {
		s.speedPulses[1] = t
	} else {
		s.speedPulses[0] = s.speedPulses[1]
		s.speedPulses[1] = t
	}
}

func (s *speedometerDev) calcSpeed() float64 {
	if s.speedPulses[0].IsZero() {
		return 0
	}
	return 0
}
