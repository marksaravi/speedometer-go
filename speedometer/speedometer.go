package speedometer

import (
	"fmt"
	"time"

	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
)

const DISPLAY_UPDATE_TIMEOUT_MS = 200

func NewSpeedometer() *speedometerDev {
	config := ReadConfigs()
	lcd := createDisplay()
	lcd.Initialise()

	speedo := speedometerDev{
		pulsePinIn:         createGpioInputPin("GPIO14"),
		resetPinIn:         createGpioInputPin("GPIO15"),
		lcd:                lcd,
		distPerPulse:       config.DistancePerPulse,
		startOfRidingTime:  time.Now(),
		resetPressedTime:   time.Now(),
		speedPulseTimeFrom: time.Time{},
		speedPulseTimeTo:   time.Time{},
		pulseCounter:       0,
		displayUpdateTurn:  0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	for {
		s.readPulse()
		s.readReset()
		ts := time.Now()
		if s.update() {
			fmt.Println(time.Since(ts))
		}

		// 	fmt.Printf("%3d, %6.2f, %6.3f, %2v\n", s.counter, s.speed, s.distance, time.Since(ts))
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
}

func (s *speedometerDev) readPulse() {
	level := s.pulsePinIn.Read()
	if s.prevPulseLevel != level && level == gpio.Low {
		s.pulseCounter++
		s.displayUpdateTime = time.Now().Add(-time.Hour)
	}
	s.prevPulseLevel = level
}

func (s *speedometerDev) calcSpeedDistanceDuration() (
	seconds, minutes, hours int, speed, distance float64,
) {
	// var speedPulseDur time.Duration = time.Second
	// if len(s.speedPulses) == 1 {
	// 	speedPulseDur = time.Since(s.speedPulses[0])
	// } else if len(s.speedPulses) == 2 {
	// 	speedPulseDur = s.speedPulses[1].Sub(s.speedPulses[0])
	// 	s.speedPulses = s.speedPulses[1:]
	// }
	// s.speed = s.distPerPulse / speedPulseDur.Seconds() * 3.6
	// if s.speed < 0.2 {
	// 	s.speed = 0
	// }
	// s.distance = s.distPerPulse * float64(s.counter)
	// s.dur = time.Since(s.startTime)

	// s.speed = randomNumber.Float64() * 50
	// s.distance = randomNumber.Float64() * 50000
	// dur = time.Duration((randomNumber.Float64() * 3600 * 8) * float64(time.Second))
	speed = 0
	seconds, minutes, hours = getSecMinHour(time.Since(s.startOfRidingTime))
	distance = s.distPerPulse * float64(s.pulseCounter)
	return
}

func getSecMinHour(d time.Duration) (int, int, int) {
	seconds := int(d.Seconds())
	return seconds % 60, seconds / 60 % 60, seconds / 3600
}

func (s *speedometerDev) update() bool {
	if time.Since(s.displayUpdateTime) < time.Millisecond*DISPLAY_UPDATE_TIMEOUT_MS {
		return false
	}
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
	s.displayUpdateTime = time.Now()
	s.displayUpdateTurn++
	if s.displayUpdateTurn == 5 {
		s.displayUpdateTurn = 0
	}
	s.lcd.UpdateDisplay()
	return true
}
