package app

import (
	"context"
	"log"
	"time"

	"github.com/marksaravi/speedometer-go/configs"
	"github.com/marksaravi/speedometer-go/models"
)

const DUR_BUFF_LEN = 20

const (
	MENU_BUTTON  = 1
	RESET_BUTTON = 2
)

type display interface {
	Initialize()
	SetInfo(speed float64, distance float64, duration time.Duration)
	Touched(x, y float64)
	ResetChannel() <-chan bool
}

type pulseSensor interface {
	Read() (bool, time.Duration)
}

type touchSensor interface {
	Touched() <-chan models.XY
}

type speedoApp struct {
	display display
	pulse   pulseSensor
	touch   touchSensor

	buttons []button
	configs configs.Configs

	pulseCounter    int64
	lastPulseTime   time.Time
	pulseDuration   time.Duration
	startTime       time.Time
	lastDisplayTime time.Time
	speed           float64
	duration        time.Duration
	distance        float64
}

func NewSpeedoApp(display display, pulse pulseSensor, touch touchSensor, configs configs.Configs) *speedoApp {
	menuButton := button{
		active:   true,
		drawable: false,
		text:     "",
		area: buttonArea{
			x1: 20,
			y1: 20,
			x2: 220,
			y2: 460,
		},
	}
	resetButton := button{
		active:   false,
		drawable: true,
		text:     "Reset",
		area: buttonArea{
			x1: 40,
			y1: 60,
			x2: 200,
			y2: 120,
		},
	}
	return &speedoApp{
		display: display,
		pulse:   pulse,
		touch:   touch,
		configs: configs,
		buttons: []button{
			resetButton,
			menuButton,
		},
	}
}

func (a *speedoApp) Start(ctx context.Context) {
	defer log.Println("Speedometer is stopped.")

	a.display.Initialize()
	a.Reset()
	for {
		select {
		case <-ctx.Done():
			return
		case <-a.display.ResetChannel():
		case xy := <-a.touch.Touched():
			a.display.Touched(xy.X, xy.Y)
		default:
			ok, dur := a.pulse.Read()
			if ok {
				a.addPulse(dur)
				// speed, distance, duration := a.calcSpeed(dur)
				// log.Printf("%6.2f, %6.2f, %v\n", speed, distance, duration)
				// if time.Since(lastDisplay) >= time.Second {
				// 	a.display.SetInfo(speed, distance, duration)
				// 	lastDisplay = time.Now()
				// }
			}
			time.Sleep(time.Millisecond)
		}
	}
}

func (a *speedoApp) Reset() {
	a.lastDisplayTime = time.Now()
	a.startTime = time.Now()
	a.lastPulseTime = time.Now().Add(-time.Second * 86400)
	a.pulseDuration = time.Since(a.lastPulseTime)
	a.pulseCounter = 0
}

func (a *speedoApp) addPulse(dur time.Duration) {
	a.lastPulseTime = time.Now()
	a.pulseCounter++
}

func (a *speedoApp) calcSpeed() (speed, distance float64, duration time.Duration) {
	speed = a.configs.DistPerPulse / float64(a.pulseDuration.Seconds()) * 3.6
	a.distance = float64(a.pulseCounter) * a.configs.DistPerPulse
	a.duration = time.Since(a.startTime)
	return
}
