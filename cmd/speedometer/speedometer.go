package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/marksaravi/drawings-go/drawings"
	"github.com/marksaravi/speedometer-go/app"
	"github.com/marksaravi/speedometer-go/configs"
	"github.com/marksaravi/speedometer-go/display"
	"github.com/marksaravi/speedometer-go/themes"
	"github.com/marksaravi/speedometer-go/touch"
	"github.com/marksaravi/speedometer-go/pulsesensor"
	"periph.io/x/host/v3"

	"github.com/marksaravi/drivers-go/colors"
	"github.com/marksaravi/drivers-go/hardware/gpio"
	"github.com/marksaravi/drivers-go/hardware/ili9341"
	"github.com/marksaravi/drivers-go/hardware/spi"
	"github.com/marksaravi/drivers-go/hardware/xpt2046"
)

func main() {
	configs := configs.ReadConfigs("./configs.yaml")
	host.Init()
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting Speedometer")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	touchSpi := spi.NewSPI(0, 0, spi.Mode0, 11, 8)
	xpt2046, err := xpt2046.NewXPT2046(ctx, &wg, touchSpi, 20)
	checkFatal(err)
	touch := touch.NewTouch(xpt2046.TouchChannel)
	lcdSpi := spi.NewSPI(1, 0, spi.Mode2, 64, 8)
	dc := gpio.NewGPIOOut("GPIO22")
	reset := gpio.NewGPIOOut("GPIO23")
	pulsePin := gpio.NewGPIOIn("GPIO17")
	ili9341, err := ili9341.NewILI9341(ili9341.LCD_320x200, lcdSpi, dc, reset)
	checkFatal(err)
	skecher := drawings.NewSketcher(ili9341, colors.BLACK)
	dis := display.NewDisplay(themes.Default, skecher, 4)
	pulse := pulsesensor.NewPulseSensor(pulsePin)
	app := app.NewSpeedoApp(dis, pulse, touch, configs)

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
