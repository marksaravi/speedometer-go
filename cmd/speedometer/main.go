package main

import (
	"log"

	"github.com/marksaravi/speedometer-go/speedometer"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	speedo := speedometer.NewSpeedometer()
	speedo.Run()
}
