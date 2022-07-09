package speedometer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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

func createDisplay() lcdDisplay {
	host.Init()
	spiConn := createSPIConnection(0, 0)
	dataCommandSelect := createGpioOutPin("GPIO22")
	reset := createGpioOutPin("GPIO23")

	ili9341Dev, err := ili9341.NewILI9341(spiConn, dataCommandSelect, reset)
	ili9341Display := display.NewRGBDisplay(ili9341Dev)
	checkFatalErr(err)
	checkFatalErr(err)
	return dashboard.NewDashboardDisplay(ili9341Display)
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

func ReadConfigs() Config {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var configs Config
	json.Unmarshal([]byte(content), &configs)
	return configs
}
