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
	Update(speed, distance float64, sec, min, hour int, secChanged, minChanged, hourChanged bool)
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
	sec          int
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
	sec, min, hour, secChanged, minChanged, hourChanged := s.getDurationChanges()
	func() {
		s.lcd.Update(s.speed, distance, sec, min, hour, secChanged, minChanged, hourChanged)
	}()
}

func (s *speedometer) getDurationChanges() (int, int, int, bool, bool, bool) {
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
	return sec, min, hour, secChanged, minChanged, hourChanged
}
