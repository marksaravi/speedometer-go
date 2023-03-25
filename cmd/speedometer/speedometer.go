package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/marksaravi/speedometer-go/app"
	"github.com/marksaravi/speedometer-go/display"
	"github.com/marksaravi/speedometer-go/speedprocessor"
	"github.com/marksaravi/speedometer-go/touch"

	"github.com/marksaravi/drivers-go/hardware/spi"
	"github.com/marksaravi/drivers-go/hardware/xpt2046"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting Speedometer")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	dis := display.NewDisplay()
	speeds := speedprocessor.NewSpeedSensor()
	touchSpi := spi.NewSPI(0, 0, spi.Mode0, 11, 8)
	xpt2046, err := xpt2046.NewXPT2046(ctx, &wg, touchSpi)
	checkFatal(err)
	touch := touch.NewTouch(xpt2046.TouchChannel)
	app := app.NewSpeedoApp(dis, speeds, touch)

	go func() {
		fmt.Println("press ENTER to stop")
		fmt.Scanln()
		cancel()
	}()

	app.Start(ctx)
	wg.Wait()
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
