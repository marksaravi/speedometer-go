package app

import (
	"context"
	"log"
	"time"

	"github.com/marksaravi/speedometer-go/models"
)

const DUR_BUFF_LEN = 20

type display interface {
	Initialize()
	SetSpeed(speed float64)
	Touched(x, y float64)
	ResetChannel() <-chan bool
}

type pulseSensor interface {
	Read() bool
}

type touchSensor interface {
	Touched() <-chan models.XY
}

type speedoApp struct {
	display display
	pulse   pulseSensor
	touch   touchSensor

	lastRead time.Time
	dts      []time.Duration
	pulses   int64
	duration time.Duration
}

func NewSpeedoApp(display display, pulse pulseSensor, touch touchSensor) *speedoApp {
	return &speedoApp{
		display: display,
		pulse:   pulse,
		touch:   touch,
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
			if a.pulse.Read() {
				a.lastRead = time.Now()
				// log.Println(speed, distance, duration)
				a.display.SetSpeed(0)
			}
		}
	}
}

func (a *speedoApp) Reset() {
	a.lastRead = time.Now().Add(time.Second*1000000)
	a.dts = make([]time.Duration, 0, DUR_BUFF_LEN)
}

func (a *speedoApp) bufferPulse() {
	dt:=time.Since(a.lastRead)
	a.dts[0]=dt
	a.lastRead = time.Now()
	a.dts = make([]time.Duration, 0, DUR_BUFF_LEN)
}
