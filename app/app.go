package app

import (
	"context"
	"log"
	"time"

	"github.com/marksaravi/speedometer-go/models"
)

type iDisplay interface {
	Initialize()
	Touched(x, y float64)
	ResetChannel() <-chan bool
}

type iSpeedProcessor interface {
	Reset()
	Update() (updated bool, speed float64, distance float64, duration time.Duration)
}

type iTouch interface {
	Touched() <-chan models.XY
}

type speedoApp struct {
	display iDisplay
	speeds  iSpeedProcessor
	touch   iTouch
}

func NewSpeedoApp(display iDisplay, speeds iSpeedProcessor, touch iTouch) *speedoApp {
	return &speedoApp{
		display: display,
		speeds:  speeds,
		touch:   touch,
	}
}

func (a *speedoApp) Start(ctx context.Context) {
	defer log.Println("Speedometer is stopped.")

	a.display.Initialize()
	a.speeds.Reset()

	for {
		select {
		case <-ctx.Done():
			return
		case <-a.display.ResetChannel():
			a.speeds.Reset()
		case xy := <-a.touch.Touched():
			a.display.Touched(xy.X, xy.Y)
		default:
			updated, speed, distance, duration := a.speeds.Update()
			if updated {
				log.Println(speed, distance, duration)
			}
		}
	}
}
