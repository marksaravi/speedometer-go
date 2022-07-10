package speedometer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/marksaravi/devices-go/devices/display"
	"github.com/marksaravi/devices-go/hardware/ili9341"
	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/spi"
)

const (
	DISPLAY_UPDATE_TIMEOUT_MS = 100
	SPEED_UPDATE_TIMEOUT      = time.Millisecond * 1303
	DIST_UPDATE_TIMEOUT       = time.Millisecond * 2521
)

func NewSpeedometer(speedPulsePin, speedResetPin gpio.PinIn, spiConn spi.Conn, dataCommandSelect gpio.PinOut, reset gpio.PinOut) *speedometerDev {
	config := ReadConfigs()
	lcd := createDisplay(spiConn, dataCommandSelect, reset)
	lcd.Initialise()

	speedo := speedometerDev{
		pulsePinIn:        speedPulsePin,
		resetPinIn:        speedResetPin,
		lcd:               lcd,
		distPerPulse:      config.DistancePerPulse,
		startOfRidingTime: time.Now(),
		resetPressedTime:  time.Now(),
		speedPulseFrom:    time.Time{},
		speedPulseTo:      time.Time{},
		pulseCounter:      0,
		prevPulseLevel:    gpio.Low,
		displayUpdateTurn: 0,
		speedLastUpdate:   time.Now(),
		distLastUpdate:    time.Now(),
		prevSecond:        100,
		prevMinute:        100,
		prevHour:          100,
	}
	return &speedo
}

func (s *speedometerDev) Run() {
	lastUpdate := time.Now()
	for {
		if s.readPulse() {
			s.pushSpeedPulse(time.Now())
			s.update()
			lastUpdate = time.Now()
		}
		if time.Since(lastUpdate) > time.Millisecond*500 {
			s.update()
			lastUpdate = time.Now()
		}
		s.readReset()
	}
}

func (s *speedometerDev) readReset() {
	if s.resetPinIn.Read() == gpio.Low {
		s.resetPressedTime = time.Now()
	}
	if time.Since(s.resetPressedTime) > time.Second*3 {
		s.reset()
	}
}

func (s *speedometerDev) reset() {
	s.pulseCounter = 0
	s.startOfRidingTime = time.Now()
	s.prevPulseLevel = gpio.Low
	s.speedPulseFrom = time.Time{}
	s.speedPulseTo = time.Time{}
}

func getSpeedPulsesZeroValue() [2]time.Time {
	return [2]time.Time{{}, {}}
}

func (s *speedometerDev) readPulse() bool {
	pulsed := false
	level := s.pulsePinIn.Read()
	if s.prevPulseLevel != level && level == gpio.Low {
		s.pulseCounter++
		pulsed = true
	}
	s.prevPulseLevel = level
	return pulsed
}

func (s *speedometerDev) calcSpeedDistanceDuration() (
	seconds, minutes, hours int, speed, distance float64,
) {

	speed = s.calcSpeed(time.Now())
	seconds, minutes, hours = getSecMinHour(time.Since(s.startOfRidingTime))
	distance = s.distPerPulse * float64(s.pulseCounter)
	return
}

func getSecMinHour(d time.Duration) (int, int, int) {
	seconds := int(d.Seconds())
	return seconds % 60, seconds / 60 % 60, seconds / 3600
}

func (s *speedometerDev) update() {
	seconds, minutes, hours, speed, distance := s.calcSpeedDistanceDuration()

	if s.prevSecond != seconds {
		s.lcd.UpdateDuration(seconds, dashboard.SECOND_CHANGED)
		s.prevSecond = seconds
	}

	if s.prevMinute != minutes {
		s.lcd.UpdateDuration(minutes, dashboard.MINUTE_CHANGED)
		s.prevMinute = minutes
	}

	if s.prevHour != hours {
		s.lcd.UpdateDuration(hours, dashboard.HOUR_CHANGED)
		s.prevHour = hours
	}

	if time.Since(s.speedLastUpdate) > SPEED_UPDATE_TIMEOUT {
		s.lcd.UpdateSpeed(speed)
		s.speedLastUpdate = time.Now()
	}

	if time.Since(s.distLastUpdate) >= DIST_UPDATE_TIMEOUT {
		s.lcd.UpdateDistance(distance)
		s.distLastUpdate = time.Now()
	}

	s.lcd.UpdateDisplay()
}

func (s *speedometerDev) pushSpeedPulse(t time.Time) {
	s.speedPulseFrom = s.speedPulseTo
	s.speedPulseTo = t
}

func (s *speedometerDev) calcSpeed(t time.Time) float64 {
	if s.speedPulseFrom.IsZero() {
		return 0
	}

	durToT1T0 := s.speedPulseTo.Sub(s.speedPulseFrom)
	durToT1 := t.Sub(s.speedPulseTo)
	dur := durToT1T0
	if durToT1 > durToT1T0 {
		dur = t.Sub(s.speedPulseFrom)
	}

	var speed float64 = 0
	if dur > time.Second*5 {
		return speed
	}
	speed = s.distPerPulse * 1000000 / float64(dur.Microseconds()) * 3.6
	return speed
}

func createDisplay(spiConn spi.Conn, dataCommandSelect, reset gpio.PinOut) lcdDisplay {

	ili9341Dev, err := ili9341.NewILI9341(spiConn, dataCommandSelect, reset)
	ili9341Display := display.NewRGBDisplay(ili9341Dev)
	checkFatalErr(err)
	checkFatalErr(err)
	return dashboard.NewDashboardDisplay(ili9341Display)
}

func checkFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadConfigs() Config {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var configs Config
	json.Unmarshal([]byte(content), &configs)
	return configs
}
