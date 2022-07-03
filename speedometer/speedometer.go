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

		speedPulseStartTime: time.Now().Add(-time.Second * 3600),
		speed:               0,
		distance:            0,
		dur:                 0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	s.speedPulseStartTime = time.Now().Add(-time.Second * 3600)
	for {
		s.pulseCounter()
		s.dur = time.Since(s.startTime)

		if time.Since(lastUpdate) >= time.Millisecond*950 {
			lastUpdate = time.Now()
			fmt.Printf("%6.2f, %6.3f, %v\n", s.speed, s.distance, s.speedPulseDur)
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

func (s *speedometerDev) pulseCounter() {
	pulse := s.input.Read()
	if pulse != s.pulse {
		if pulse == gpio.Low {
			s.counter++
			if s.counter > 1 {
				s.speedPulseStartTime = s.lastPulse
			}
			s.lastPulse = time.Now()
		}
		s.pulse = pulse
	}
	speedPulseDur := time.Since(s.speedPulseStartTime)
	s.speed = s.distPerPulse / speedPulseDur.Seconds() * 3.6
	if s.speed < 0.25 {
		s.speed = 0
	}
	s.distance = s.distPerPulse * float64(s.counter)
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
