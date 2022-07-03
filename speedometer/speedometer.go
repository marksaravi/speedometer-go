package speedometer

import (
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
		pulseTime:  time.Now(),
		pulseDur:   0,
		resetLevel: gpio.Low,
		resetTime:  time.Now(),

		speed:    0,
		distance: 0,
		dur:      0,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	for {
		speed, distance, pulsed := s.readPulse()
		s.dur = time.Since(s.startTime)
		if pulsed {
			s.speed = speed
			s.distance = distance
		}
		if time.Since(lastUpdate) >= time.Millisecond*900 {
			s.update()
			lastUpdate = time.Now()
			if pulsed {
				time.Sleep(time.Microsecond * time.Duration(s.sleepAfterPulseMS))
			}
		}
		s.readReset()
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
