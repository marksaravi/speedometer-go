package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		time.Sleep(time.Millisecond * 500)

		for i := 0; i < 4; i++ {
			led.High()
			time.Sleep(time.Millisecond * 150)
			led.Low()
			time.Sleep(time.Millisecond * 150)
		}
	}
}
