package main

import (
	"time"
)

type lcdDisplay interface {
	Initialise()
	Update(speed, distance float64, duration time.Duration)
}

func main() {
	lcd := createDisplay()
	lcd.Initialise()
	process(lcd)
}

func process(lcd lcdDisplay) {
	lcd.Update(17.3, 847.45, time.Second*4325)
}
