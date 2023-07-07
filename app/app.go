package app

import (
	"context"
	"log"
	"time"

	"github.com/marksaravi/speedometer-go/configs"
	"github.com/marksaravi/speedometer-go/models"
)

const DUR_BUFF_LEN = 20

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

	configs configs.Configs

	durations       []time.Duration
	pulses    int64
	startTime time.Time
}

func NewSpeedoApp(display display, pulse pulseSensor, touch touchSensor, configs configs.Configs) *speedoApp {
	return &speedoApp{
		display: display,
		pulse:   pulse,
		touch:   touch,
		configs: configs,
	}
}

func (a *speedoApp) Start(ctx context.Context) {
	defer log.Println("Speedometer is stopped.")

	a.display.Initialize()
	a.Reset()
	lastDisplay := time.Now()
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
				speed, distance, duration := a.calcSpeed(dur)
				if time.Since(lastDisplay)>=time.Second {
					a.display.SetInfo(speed, distance, duration)
					log.Printf("%6.2f, %6.2f, %v\n", speed, distance, duration)
					lastDisplay = time.Now()
				}
			}
		}
	}
}

func (a *speedoApp) Reset() {
	a.durations = make([]time.Duration, DUR_BUFF_LEN)
	for i:=0; i<DUR_BUFF_LEN; i++ {
		a.durations[i]=time.Second*86400
	}
	a.startTime = time.Now()
}

func (a *speedoApp) calcSpeed(dur time.Duration) (speed, distance float64, duration time.Duration) {
	for i:=1; i<DUR_BUFF_LEN; i++ {
		a.durations[i]=a.durations[i-1]
	}
	a.durations[0]=dur
	a.pulses++
	speed = a.configs.DistPerPulse / float64(dur.Seconds()) * 3.6
	distance = float64(a.pulses)*a.configs.DistPerPulse
	duration = time.Since(a.startTime)
	return
}
