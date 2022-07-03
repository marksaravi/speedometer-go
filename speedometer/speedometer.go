package speedometer

import (
	"fmt"
	"math"
	"time"

	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
)

func NewSpeedometer() *speedometerDev {
	config := ReadConfigs()
	lcd := createDisplay()
	lcd.Initialise()

	speedo := speedometerDev{
		input: createGpioInputPin("GPIO14"),
		reset: createGpioInputPin("GPIO15"),
		lcd:   lcd,

		distPerPulse:      config.DistancePerPulse,
		sleepAfterPulseMS: config.SleepAfterPulseMS,

		startTime:  time.Now(),
		counter:    0,
		pulse:      gpio.Low,
		pulseTime:  time.Now(),
		pulseDur:   0,
		resetLevel: gpio.Low,
		resetTime:  time.Now(),

		speed:    0,
		distance: 0,
		sec:      0,
		min:      0,
		hour:     0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	loops := 0
	for {
		loops++
		speed, distance, changed := s.readPulse()
		s.readReset()

		if time.Since(lastUpdate) >= time.Second {
			lastUpdate = time.Now()
			fmt.Println(s.counter, speed, distance, loops)
			loops = 0
			s.update(speed, distance, changed)
		}
		if changed {
			time.Sleep(time.Microsecond * time.Duration(s.sleepAfterPulseMS))
		}
	}
}

func (s *speedometerDev) resetAll() {
	s.counter = 0
	s.startTime = time.Now()
	s.pulseTime = time.Now()
	s.pulse = gpio.Low
	s.resetLevel = gpio.Low
}

func (s *speedometerDev) readPulse() (float64, float64, bool) {
	pulse := s.input.Read()
	var speed float64 = 0
	changed := false
	if pulse != s.pulse {
		if pulse == gpio.Low {
			s.counter++
			s.pulseDur = time.Since(s.pulseTime)
			s.pulseTime = time.Now()
			speed = s.distPerPulse / s.pulseDur.Seconds() * 1000 / 3600
			changed = true
		}
		s.pulse = pulse
	}
	distance := float64(s.counter) * s.distPerPulse
	return speed, distance, changed
}

func (s *speedometerDev) readReset() {
	reset := s.reset.Read()
	if reset == gpio.Low {
		s.resetTime = time.Now()
	}
	if reset == gpio.High {
		if time.Since(s.resetTime) > time.Second*3 {
			s.resetAll()
		}
	}
}

func (s *speedometerDev) getDurationChanges() (bool, bool, bool) {
	dur := time.Since(s.startTime)
	sec := int(dur.Seconds()) % 60
	min := sec / 60 % 60
	hour := sec / 3600

	secChanged := sec != s.sec
	s.sec = sec
	minChanged := min != s.min
	s.min = min
	hourChanged := hour != s.hour
	s.hour = hour
	return secChanged, minChanged, hourChanged
}

func (s *speedometerDev) updateSpeed(speed float64) bool {
	if math.Abs(speed-s.speed) >= MIN_SPEED_UPDATE {
		s.speed = speed
		return true
	}
	return false
}

func (s *speedometerDev) updateDistance(distance float64) bool {
	if math.Abs(distance-s.distance) >= MIN_DISTANCE_UPDATE {
		s.distance = distance
		return true
	}
	return false
}

func (s *speedometerDev) update(speed, distance float64, changed bool) {
	secChanged, minChanged, hourChanged := s.getDurationChanges()
	speedChanged := s.updateSpeed(speed)
	distanceChanged := s.updateDistance(distance)
	func() {
		s.lcd.Update(dashboard.DisplayData{
			Speed:           s.speed,
			SpeedChanged:    speedChanged,
			Distance:        s.distance,
			DistanceChanged: distanceChanged,
			Sec:             s.sec,
			SecChanged:      secChanged,
			Min:             s.min,
			MinChanged:      minChanged,
			Hour:            s.hour,
			HourChanged:     hourChanged,
		})
	}()
}
