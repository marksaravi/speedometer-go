package speedometer

import (
	"fmt"
	"time"

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
		resetLevel: gpio.Low,
		resetTime:  time.Now(),

		speedPulses: make([]time.Time, 0),
		speed:       0,
		distance:    0,
		dur:         0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	ts := time.Now()
	var max time.Duration = 0
	s.speedPulses = append(s.speedPulses, time.Now().Add(-time.Second*86400))
	for {
		if d := time.Since(ts); d > max {
			max = d
		}
		ts = time.Now()
		s.pulseCounter()

		if time.Since(lastUpdate) >= time.Millisecond*950 {
			s.updateSpeedDistanceDuration()
			lastUpdate = time.Now()
			fmt.Printf("%3d, %6.2f, %6.3f, %3d, %2d\n", s.counter, s.speed, s.distance, max.Milliseconds(), len(s.speedPulses))
			max = 0
			s.update()
		}
	}
}

func (s *speedometerDev) resetAll() {
	s.counter = 0
	s.startTime = time.Now()
	s.pulse = gpio.Low
	s.resetLevel = gpio.Low
}

var tpulse = time.Now()
var MILLISECONDS float64 = 64

func (s *speedometerDev) pulseFaker() gpio.Level {
	pulse := s.pulse
	if s.pulse == gpio.Low && time.Since(tpulse) >= time.Millisecond*time.Duration((MILLISECONDS)) {
		pulse = gpio.High
		tpulse = time.Now()
	}

	if s.pulse == gpio.High && time.Since(tpulse) >= time.Millisecond*time.Duration((MILLISECONDS)/8) {
		pulse = gpio.Low
		tpulse = time.Now()
	}
	return pulse
}

func (s *speedometerDev) pulseCounter() bool {
	// pulse := s.input.Read()
	pulse := s.pulseFaker()
	isPulsed := false

	if pulse != s.pulse {
		if pulse == gpio.Low {
			s.counter++
			if len(s.speedPulses) == 2 {
				s.speedPulses = s.speedPulses[1:]
			}
			s.speedPulses = append(s.speedPulses, time.Now())
			isPulsed = true
		}
		s.pulse = pulse
	}
	return isPulsed
}

func (s *speedometerDev) updateSpeedDistanceDuration() {
	var speedPulseDur time.Duration = time.Second
	if len(s.speedPulses) == 1 {
		speedPulseDur = time.Since(s.speedPulses[0])
	} else if len(s.speedPulses) == 2 {
		speedPulseDur = s.speedPulses[1].Sub(s.speedPulses[0])
		s.speedPulses = s.speedPulses[1:]
	}
	s.speed = s.distPerPulse / speedPulseDur.Seconds() * 3.6
	if s.speed < 0.2 {
		s.speed = 0
	}

	s.distance = s.distPerPulse * float64(s.counter)
	s.dur = time.Since(s.startTime)
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

func getSecMinHour(d time.Duration) (int, int, int) {
	seconds := int(d.Seconds()) % 60
	return seconds, seconds / 60 % 60, seconds / 3600
}

func (s *speedometerDev) update() {
	seconds, minutes, hours := getSecMinHour(s.dur)
	s.lcd.UpdateSecond(seconds)
	s.lcd.UpdateMinute(minutes)
	s.lcd.UpdateHour(hours)
	s.lcd.UpdateSpeed(s.speed)
	s.lcd.UpdateDistance(s.distance)
	func() {
		s.lcd.UpdateDisplay()
	}()
}
