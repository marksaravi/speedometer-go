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

		pulsePerSecond:    0,
		pulseCountingTime: time.Now(),
		speed:             0,
		distance:          0,
		dur:               0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	s.pulsePerSecond = 0
	s.pulseCountingTime = time.Now().Add(-time.Second * 3600)
	for {
		s.pulseCounter()

		if time.Since(lastUpdate) >= time.Millisecond*950 {
			lastUpdate = time.Now()
			fmt.Println(s.speed, s.distance, s.dur)
			// s.update()
			// 	if pulsed {
			// 		time.Sleep(time.Millisecond * time.Duration(s.sleepAfterPulseMS))
			// 	}
		}
		// s.readReset()
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
	pulsed := false
	if pulse != s.pulse {
		if pulse == gpio.Low {
			pulsed = true
		}
		s.pulse = pulse
	}
	if pulsed {
		s.counter++
		s.pulsePerSecond++
	}
	if time.Since(s.pulseCountingTime) >= time.Second {
		fmt.Println(s.pulsePerSecond)
		s.pulseCountingTime = time.Now()
		if s.pulsePerSecond == 0 {
			s.speed = 0
		} else {

		}
		s.pulsePerSecond = 0
	}
}

func (s *speedometerDev) updateSpeedDistDur() {
	// // s.speed = s.distPerPulse / s.pulseDur.Seconds() * 3600 / 1000
	// s.speed = s.pulsePerSecond
	// s.distance = float64(s.counter) * s.distPerPulse
	// s.dur = time.Since(s.startTime)
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
	// func() {
	// 	s.lcd.UpdateDisplay()
	// }()
}
