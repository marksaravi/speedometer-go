package touchdisplay

import "github.com/marksaravi/drivers-go/hardware/spi"

type touchDisplay struct {
	lcd spi.SPI
}
