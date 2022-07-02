package main

import (
	"fmt"
	"math"
	"time"

	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
)

type Config struct {
	DistancePerPulse float64 `json:"distance-per-pulse"`
}

type lcdDisplay interface {
	Initialise()
	Update(dashboard.DisplayData)
}

type speedometer struct {
	counter   int64
	startTime time.Time

	pulse     gpio.Level
	pulseTime time.Time
	pulseDur  time.Duration

	resetLevel gpio.Level
	resetTime  time.Time

	distPerPulse float64
	sec          int
	min          int
	hour         int
	distance     float64
	speed        float64

	input gpio.PinIn
	reset gpio.PinIn
	lcd   lcdDisplay
}

func main() {
	config := ReadConfigs()
	lcd := createDisplay()
	lcd.Initialise()

	speedo := speedometer{
		startTime:  time.Now(),
		counter:    0,
		pulse:      gpio.Low,
		pulseTime:  time.Now(),
		pulseDur:   0,
		resetLevel: gpio.Low,
		resetTime:  time.Now(),

		distPerPulse: config.DistancePerPulse,
		speed:        0,
		distance:     0,
		sec:          0,
		min:          0,
		hour:         0,

		input: createGpioInputPin("GPIO14"),
		reset: createGpioInputPin("GPIO15"),
		lcd:   lcd,
	}
	speedo.process()
}

func (s *speedometer) process() {
	lastUpdate := time.Now()
	for {
		speed, distance, changed := s.readPulse()
		s.readReset()
		time.Sleep(time.Millisecond)
		if time.Since(lastUpdate) >= time.Second {
			lastUpdate = time.Now()
			fmt.Println(s.counter, speed, distance)
			s.update(speed, distance, changed)
		}

	}
}

func (s *speedometer) resetAll() {
	s.counter = 0
	s.startTime = time.Now()
	s.pulseTime = time.Now()
	s.pulse = gpio.Low
	s.resetLevel = gpio.Low
}

func (s *speedometer) readPulse() (float64, float64, bool) {
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

func (s *speedometer) readReset() {
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

func (s *speedometer) getDurationChanges() (bool, bool, bool) {
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

func (s *speedometer) updateSpeed(speed float64) bool {
	if math.Abs(speed-s.speed) >= dashboard.SPEED_RESOLUTION {
		s.speed = speed
		return true
	}
	return false
}

func (s *speedometer) updateDistance(distance float64) bool {
	if math.Abs(distance-s.distance) >= dashboard.DISTANCE_RESOLUTION {
		s.distance = distance
		return true
	}
	return false
}

func (s *speedometer) update(speed, distance float64, changed bool) {
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
