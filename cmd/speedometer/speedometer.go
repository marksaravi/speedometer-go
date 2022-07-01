package main

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
)

type lcdDisplay interface {
	Initialise()
	Update(speed, distance float64, duration time.Duration)
}

func main() {
	lcd := createDisplay()
	lcd.Initialise()
	input := createGpioInputPin("GPIO6")
	process(lcd, input)
}

func process(lcd lcdDisplay, input gpio.PinIn) {
	const PERIMETER float64 = 2.2
	const PULSE_PER_PERIMETER = 8
	const DIST_PER_PULSE = PERIMETER / PULSE_PER_PERIMETER

	start := time.Now()
	var counter int = 0
	ts := time.Now()
	prevTime := time.Now()
	currTime := time.Now()
	for {
		time.Sleep(time.Millisecond * 10)
		input.Read()
		counter++
		currTime = time.Now()
		dt := currTime.Sub(prevTime)
		prevTime = currTime
		speed := DIST_PER_PULSE / dt.Seconds() * 1000 / 3600
		dur := time.Since(start)
		distance := DIST_PER_PULSE * float64(counter)

		if time.Since(ts) >= time.Second {
			fmt.Println(counter)
			ts = time.Now()
			func() {
				lcd.Update(speed, distance, dur)
			}()
		}
	}
}
