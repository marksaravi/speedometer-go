package main

import (
	"context"
	"fmt"
	"log"
	"speedometer-go/app"
	"speedometer-go/display"
	"speedometer-go/speedprocessor"
	"speedometer-go/touch"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting Speedometer")
	dis := display.NewDisplay()
	speeds := speedprocessor.NewSpeedSensor()
	touch := touch.NewTouch()
	app := app.NewSpeedoApp(dis, speeds, touch)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("press ENTER to stop")
		fmt.Scanln()
		cancel()
	}()

	app.Start(ctx)
}
