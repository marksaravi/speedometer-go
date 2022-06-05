package main

import (
	"fmt"
	"log"
	"time"

	"github.com/marksaravi/devices-go/devices/display"
	"github.com/marksaravi/devices-go/hardware/ili9341"
	"github.com/marksaravi/speedometer-go/dashboard"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/sysfs"
)

func main() {
	_, _ = createDisplay()
	time.Sleep(time.Second)
	// counter := 0
	// for {
	// 	time.Sleep(time.Second / 8)
	// 	pulse <- time.Now()
	// 	if counter == 100 {
	// 		counter = 0
	// 		reset <- true
	// 	}
	// 	counter++
	// }
}

func createDisplay() (chan<- time.Time, chan<- bool) {
	host.Init()
	spiConn := createSPIConnection(0, 0)
	dataCommandSelect := createGpioOutPin("GPIO22")
	reset := createGpioOutPin("GPIO23")
	var display display.RGB565Display
	var err error
	display = ili9341.NewILI9341(spiConn, dataCommandSelect, reset)
	checkFatalErr(err)
	return dashboard.NewDashboardDisplay(display)
}

func createGpioOutPin(gpioPinNum string) gpio.PinOut {
	var pin gpio.PinOut = gpioreg.ByName(gpioPinNum)
	if pin == nil {
		checkFatalErr(fmt.Errorf("failed to create GPIO pin %s", gpioPinNum))
	}
	pin.Out(gpio.Low)
	return pin
}

func createSPIConnection(busNumber int, chipSelect int) spi.Conn {
	spibus, _ := sysfs.NewSPI(
		busNumber,
		chipSelect,
	)
	spiConn, err := spibus.Connect(
		physic.Frequency(12)*physic.MegaHertz,
		spi.Mode3,
		8,
	)
	checkFatalErr(err)
	return spiConn
}
func checkFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
