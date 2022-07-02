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

	input gpio.PinIn
	reset gpio.PinIn
	lcd   lcdDisplay
}

func main() {
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

		distPerPulse: 0.275,
		speed:        0,

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

// func process(lcd lcdDisplay, input gpio.PinIn, reset gpio.PinIn) {
// 	const PERIMETER float64 = 2.2
// 	const PULSE_PER_PERIMETER = 8
// 	const DIST_PER_PULSE = PERIMETER / PULSE_PER_PERIMETER

// 	start := time.Now()
// 	var counter int = 0
// 	var prveLevel gpio.Level = gpio.Low
// 	var prevLevelTime time.Time = time.Now()
// 	var speed float64 = 0
// 	var distance float64 = 0

// 	for {
// 		level := input.Read()
// 		if level != prveLevel {
// 			level = prveLevel
// 			if level == gpio.High {
// 				counter++
// 				dt := time.Since(prevLevelTime)
// 				prevLevelTime = time.Now()
// 			}
// 		}

// 		// currTime = time.Now()
// 		// dt := currTime.Sub(prevTime)
// 		// prevTime = currTime
// 		// speed := DIST_PER_PULSE / dt.Seconds() * 1000 / 3600
// 		// dur := time.Since(start)
// 		// distance := DIST_PER_PULSE * float64(counter)

// 		// if time.Since(ts) >= time.Second {
// 		// 	fmt.Println(level)
// 		// 	fmt.Println(counter)
// 		// 	ts = time.Now()
// 		// 	func() {
// 		// 		lcd.Update(speed, distance, dur)
// 		// 	}()
// 		// }
// 	}
// }
