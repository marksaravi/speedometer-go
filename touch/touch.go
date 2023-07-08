package touch

import (
	"github.com/marksaravi/speedometer-go/models"
)

type touch struct {
	touched <-chan models.XY
}

func NewTouch(touched chan models.XY) *touch {
	return &touch{
		touched: touched,
	}
}

func (t *touch) Touched() <-chan models.XY {
	return t.touched
}

func ConverionFactors(x1t, x2t, x1s, x2s float64) (float64, float64) {

	a := (x2s-x1s)/(x2t-x1t)
	b := x1s - a*x1t
	return a, b
}

func Convert(xt,yt, ax, bx, ay, by float64 ) (float64, float64) {
	xs := xt * ax + bx
	ys := yt * ay + by
	return xs, ys
}