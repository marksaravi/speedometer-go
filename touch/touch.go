package touch

import (
	"github.com/marksaravi/speedometer-go/models"
)

type touch struct {
	touched <-chan models.XY
	ax, bx, ay, by float64
}

func NewTouch(touched chan models.XY) *touch {
	//512, 1600        1600, 1008
	//40, 100          200, 160
	ax, bx := ConverionFactors(512, 1600, 40, 200)
	ay, by := ConverionFactors(1600, 1008, 100, 160)
	return &touch{
		touched: touched,
		ax: ax,
		bx: bx,
		ay: ay,
		by: by,
	}
}

func (t *touch) Touched() <-chan models.XY {
	return t.touched
}

func (t *touch) TouchConvert(xy models.XY) (float64, float64) {
	x, y := Convert(xy.X, xy.Y, t.ax, t.bx, t.ay, t.by)
	return x, y
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