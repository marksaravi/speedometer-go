package speedometer

import (
	"math"
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
		changed:  false,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	for {
		speed, distance, pulsed := s.readPulse()
		if pulsed {
			s.updateSpeed(speed)
			s.updateDistance(distance)
		}
		if time.Since(lastUpdate) >= time.Second {
			s.updateDuration()
			s.update()
		}
		if pulsed {
			time.Sleep(time.Microsecond * time.Duration(s.sleepAfterPulseMS))
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

func (s *speedometerDev) updateDuration() {
	dur := time.Since(s.startTime)
	seconds, minutes, hours := getSecMinHour(dur)
	prevSeconds, prevMinutes, prevHours := getSecMinHour(s.dur)
	s.dur = dur
	s.changed = seconds != prevSeconds || minutes != prevMinutes || hours != prevHours
}

func (s *speedometerDev) updateSpeed(speed float64) {
	if math.Abs(speed-s.speed) >= MIN_SPEED_UPDATE {
		s.speed = speed
		s.changed = true
	}
}

func (s *speedometerDev) updateDistance(distance float64) {
	if math.Abs(distance-s.distance) >= MIN_DISTANCE_UPDATE {
		s.distance = distance
		s.changed = true
	}
}

func (s *speedometerDev) update() {
	func() {
		if s.changed {
			seconds, minutes, hours := getSecMinHour(s.dur)
			s.lcd.UpdateSecond(seconds)
			s.lcd.UpdateSecond(minutes)
			s.lcd.UpdateSecond(hours)
			s.lcd.UpdateSpeed(s.speed)
			s.lcd.UpdateDistance(s.distance)
			s.lcd.UpdateDisplay()
			s.changed = false
		}
	}()
}
