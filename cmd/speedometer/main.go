package main

import (
	"fmt"
	"log"

	"github.com/marksaravi/speedometer-go/speedometer"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/sysfs"
)

func main() {
	host.Init()
	spiConn := createSPIConnection(0, 0)
	dataCommandSelect := createGpioOutPin("GPIO22")
	reset := createGpioOutPin("GPIO23")
	speedPulsePinIn := createGpioInputPin("GPIO14")
	speedResetPinIn := createGpioInputPin("GPIO15")
	log.SetFlags(log.Lmicroseconds)
	speedo := speedometer.NewSpeedometer(
		speedPulsePinIn,
		speedResetPinIn,
		spiConn,
		dataCommandSelect,
		reset,
	)
	speedo.Run()
}

func createGpioOutPin(gpioPinNum string) gpio.PinOut {
	var pin gpio.PinOut = gpioreg.ByName(gpioPinNum)
	if pin == nil {
		checkFatalErr(fmt.Errorf("failed to create GPIO pin %s", gpioPinNum))
	}
	pin.Out(gpio.Low)
	return pin
}

func createGpioInputPin(gpioPinNum string) gpio.PinIn {
	var pin gpio.PinIn = gpioreg.ByName(gpioPinNum)
	if pin == nil {
		log.Fatal(fmt.Errorf("failed to create GPIO pin %s", gpioPinNum))
	}
	if err := pin.In(gpio.PullDown, gpio.RisingEdge); err != nil {
		log.Fatal(err)
	}
	return pin
}

func createSPIConnection(busNumber int, chipSelect int) spi.Conn {
	spibus, _ := sysfs.NewSPI(
		busNumber,
		chipSelect,
	)
	spiConn, err := spibus.Connect(
		physic.Frequency(48)*physic.MegaHertz,
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
