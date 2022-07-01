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
	lcd.Update(17.3, 847.45, time.Second*4325)
	ts := time.Now()
	var counter int = 0
	for {
		input.Read()
		counter++
		if time.Since(ts) >= time.Second {
			fmt.Println(counter)
			counter = 0
			ts = time.Now()
		}
	}
}
