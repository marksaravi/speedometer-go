package main

import (
	"time"

	"periph.io/x/conn/v3/gpio"
)

type Config struct {
	DistancePerPulse float64 `json:"distance-per-pulse"`
}

type lcdDisplay interface {
	Initialise()
	Update(speed, distance float64, duration time.Duration)
}

type speedometer struct {
	counter   int64
	startTime time.Time

	prevPulse     gpio.Level
	prevPulseTime time.Time
	pulseDur      time.Duration

	prevReset gpio.Level
	resetTime time.Time

	distPerPulse float64
	speed        float64
	min          int
	hour         int

	input gpio.PinIn
	reset gpio.PinIn
	lcd   lcdDisplay
}

func main() {
	config := ReadConfigs()
	lcd := createDisplay()
	lcd.Initialise()

	speedo := speedometer{
		startTime:     time.Now(),
		counter:       0,
		prevPulse:     gpio.Low,
		prevPulseTime: time.Now(),
		pulseDur:      0,
		prevReset:     gpio.Low,
		resetTime:     time.Now(),

		distPerPulse: config.DistancePerPulse,
		speed:        0,
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
		s.readPulse()
		s.readReset()
		time.Sleep(time.Millisecond)
		if time.Since(lastUpdate) >= time.Second {
			lastUpdate = time.Now()
			s.update()
		}

	}
}

func (s *speedometer) resetAll() {
	s.counter = 0
	s.startTime = time.Now()
	s.prevPulseTime = time.Now()
	s.prevPulse = gpio.Low
	s.prevReset = gpio.Low
}

func (s *speedometer) readPulse() {
	pulse := s.input.Read()
	if pulse != s.prevPulse {
		if pulse == gpio.High {
			s.counter++
			s.pulseDur = time.Since(s.prevPulseTime)
			s.prevPulseTime = time.Now()
			s.speed = s.distPerPulse / s.pulseDur.Seconds() * 1000 / 3600
		}
		s.prevPulse = pulse
	}
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

func (s *speedometer) update() {
	distance := s.distPerPulse * float64(s.counter)
	dur := time.Since(s.startTime)
	func() {
		s.lcd.Update(s.speed, distance, dur)
	}()
}
